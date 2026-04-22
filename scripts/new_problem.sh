#!/usr/bin/env bash
# Usage:
#   ./scripts/new_problem.sh "11. Container with most water" "https://leetcode.com/problems/container-with-most-water/description/"
#   ./scripts/new_problem.sh --url "https://leetcode.com/problems/validate-binary-search-tree/description/"
#   ./scripts/new_problem.sh "https://leetcode.com/problems/two-sum/"   # single LeetCode URL is auto-detected
# Run from repo root. Creates new/YYYYMMDD/SLUG.md and src/SLUG/main.go (Monday = first day of current week).
# URL mode fetches title, statement, and code template via https://leetcode.com/graphql (no browser).
# Idempotent: skips if files already exist.

set -e

usage() {
  echo "Usage:" >&2
  echo "  $0 \"<title>\" \"<leetcode-url>\"" >&2
  echo "  $0 --url \"<leetcode-url>\"" >&2
  echo "  $0 \"<leetcode-url>\"   # if the only arg looks like a LeetCode problem URL" >&2
  exit 1
}

# Extract title slug from https://leetcode.com/problems/<slug>/...
leetcode_slug_from_url() {
  local url="$1"
  if [[ "$url" =~ leetcode\.com/problems/([^/\?#]+) ]]; then
    printf '%s' "${BASH_REMATCH[1]}"
  else
    echo "Could not parse LeetCode problem slug from URL (expected .../problems/<slug>/...)" >&2
    return 1
  fi
}

# Fetch problem metadata + HTML statement + Go template; print one JSON object to stdout.
fetch_leetcode_json() {
  local slug="$1"
  python3 - "$slug" <<'PY'
import html
import json
import re
import sys
import urllib.error
import urllib.request

slug = sys.argv[1]
query = """query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId
    title
    content
    difficulty
    codeSnippets {
      lang
      langSlug
      code
    }
  }
}"""
payload = json.dumps({"query": query, "variables": {"titleSlug": slug}}).encode("utf-8")
req = urllib.request.Request(
    "https://leetcode.com/graphql",
    data=payload,
    headers={
        "Content-Type": "application/json",
        "Referer": "https://leetcode.com/",
        "User-Agent": "Mozilla/5.0 (compatible; DnA-new-problem/1.0)",
    },
    method="POST",
)
try:
    with urllib.request.urlopen(req, timeout=45) as r:
        resp = json.load(r)
except urllib.error.URLError as e:
    print("Network error talking to LeetCode GraphQL:", e, file=sys.stderr)
    sys.exit(1)

q = (resp.get("data") or {}).get("question")
if not q:
    print("Unknown titleSlug or empty response from LeetCode:", slug, file=sys.stderr)
    if resp.get("errors"):
        print(json.dumps(resp["errors"], indent=2), file=sys.stderr)
    sys.exit(1)

raw_html = q.get("content") or ""
# Rough HTML -> plain text for notes (readable offline).
s = html.unescape(raw_html)
s = s.replace("\xa0", " ")
s = re.sub(r"(?i)<\s*br\s*/?>", "\n", s)
s = re.sub(r"(?i)</(p|div|h[1-4]|blockquote)\s*>", "\n\n", s)
s = re.sub(r"(?i)</(li|tr)\s*>", "\n", s)
s = re.sub(r"(?i)</pre\s*>", "\n\n", s)
s = re.sub(r"<[^>]+>", " ", s)
s = re.sub(r"[ \t\f\v]+", " ", s)
s = re.sub(r"\n[ \t]*\n[ \t]*", "\n\n", s)
text = re.sub(r"\n{3,}", "\n\n", s.strip())

out = {
    "title": f'{q["questionFrontendId"]}. {q["title"]}',
    "url": f"https://leetcode.com/problems/{slug}/description/",
    "difficulty": q.get("difficulty") or "",
    "problem_plain": text,
    "go_template": "",
}

for snip in q.get("codeSnippets") or []:
    if (snip.get("langSlug") or "").lower() == "golang":
        out["go_template"] = snip.get("code") or ""
        break

json.dump(out, sys.stdout, ensure_ascii=False)
sys.stdout.write("\n")
PY
}

TITLE=""
URL=""
FETCHED_JSON=""

if [[ "${1:-}" == "--url" || "${1:-}" == "-u" ]]; then
  [[ -n "${2:-}" ]] || usage
  URL="$2"
  SLUG_RAW=$(leetcode_slug_from_url "$URL") || exit 1
  FETCHED_JSON=$(fetch_leetcode_json "$SLUG_RAW")
  TITLE=$(python3 -c "import json,sys; print(json.load(sys.stdin)['title'])" <<<"$FETCHED_JSON")
  URL=$(python3 -c "import json,sys; print(json.load(sys.stdin)['url'])" <<<"$FETCHED_JSON")
elif [[ $# -eq 1 && "$1" =~ ^https?://leetcode\.com/problems/ ]]; then
  URL="$1"
  SLUG_RAW=$(leetcode_slug_from_url "$URL") || exit 1
  FETCHED_JSON=$(fetch_leetcode_json "$SLUG_RAW")
  TITLE=$(python3 -c "import json,sys; print(json.load(sys.stdin)['title'])" <<<"$FETCHED_JSON")
  URL=$(python3 -c "import json,sys; print(json.load(sys.stdin)['url'])" <<<"$FETCHED_JSON")
elif [[ $# -eq 2 ]]; then
  TITLE="${1:?Usage: $0 \"<title>\" \"<url>\"}"
  URL="${2:?Usage: $0 \"<title>\" \"<url>\"}"
else
  usage
fi

# Monday of current week (YYYYMMDD), portable (Linux + macOS)
MONDAY=$(python3 -c "from datetime import datetime, timedelta; d=datetime.now(); m=d-timedelta(days=d.weekday()); print(m.strftime('%Y%m%d'))")
TODAY=$(date +%Y%m%d)

# Slug from title: "11. Container with most water" -> 11_container_with_most_water
if [[ "$TITLE" =~ ^([0-9]+)\.\s*(.*)$ ]]; then
  NUM="${BASH_REMATCH[1]}"
  REST="${BASH_REMATCH[2]}"
  REST_CLEAN=$(echo "$REST" | tr 'A-Z' 'a-z' | sed 's/[ .]\+/_/g' | sed 's/_\+/_/g' | sed 's/^_//;s/_$//')
  SLUG="${NUM}_${REST_CLEAN}"
else
  SLUG=$(echo "$TITLE" | tr 'A-Z' 'a-z' | sed 's/[ .]\+/_/g' | sed 's/_\+/_/g' | sed 's/^_//;s/_$//')
fi

NEW_DIR="new/${MONDAY}"
MD_FILE="${NEW_DIR}/${SLUG}.md"
SRC_DIR="src/${SLUG}"
MAIN_GO="${SRC_DIR}/main.go"

DIFFICULTY=""
PROBLEM_PLAIN=""
GO_TEMPLATE=""
if [[ -n "$FETCHED_JSON" ]]; then
  DIFFICULTY=$(python3 -c "import json,sys; print(json.load(sys.stdin).get('difficulty',''))" <<<"$FETCHED_JSON")
  PROBLEM_PLAIN=$(python3 -c "import json,sys; print(json.load(sys.stdin).get('problem_plain',''))" <<<"$FETCHED_JSON")
  GO_TEMPLATE=$(python3 -c "import json,sys; print(json.load(sys.stdin).get('go_template',''))" <<<"$FETCHED_JSON")
fi

# Create note (idempotent)
mkdir -p "$NEW_DIR"
if [[ -f "$MD_FILE" ]]; then
  echo "Skipped (exists): $MD_FILE"
else
  if [[ -n "$FETCHED_JSON" ]]; then
    {
      cat <<EOF
# ${TITLE}

## Link

[Leetcode](${URL})

## Difficulty

${DIFFICULTY}

## Problem

EOF
      printf '%s\n' "$PROBLEM_PLAIN"
      cat <<EOF

## Solve Date

${TODAY}

## Solution

[Code](../../${SRC_DIR}/main.go)

## Status

<!-- NEW | IN PROGRESS | DONE -->

## Core idea

<!-- Pattern, key insight, time/space -->

## Failure

<!-- What went wrong / what to avoid next time -->

## Success

<!-- What went well / what to reuse -->

## Tags

<!-- e.g. array, hash-set, sorting -->
EOF
    } >"$MD_FILE"
  else
    cat >"$MD_FILE" <<EOF
# ${TITLE}

## Link

[Leetcode](${URL})

## Solve Date

${TODAY}

## Solution

[Code](../../${SRC_DIR}/main.go)

## Status

<!-- NEW | IN PROGRESS | DONE -->

## Core idea

<!-- Pattern, key insight, time/space -->

## Failure

<!-- What went wrong / what to avoid next time -->

## Success

<!-- What went well / what to reuse -->

## Tags

<!-- e.g. array, hash-set, sorting -->
EOF
  fi
  echo "Created: $MD_FILE"
fi

# Create src folder and main.go (idempotent)
if [[ -f "$MAIN_GO" ]]; then
  echo "Skipped (exists): $MAIN_GO"
else
  mkdir -p "$SRC_DIR"
  if [[ -n "$GO_TEMPLATE" ]]; then
    printf '%s\n' "$GO_TEMPLATE" >"$MAIN_GO"
  else
    cat >"$MAIN_GO" <<'GOEOF'
package main

func main() {
}
GOEOF
  fi
  echo "Created: $MAIN_GO"
fi
