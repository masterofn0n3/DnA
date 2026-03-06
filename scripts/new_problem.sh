#!/usr/bin/env bash
# Usage: ./scripts/new_problem.sh "11. Container with most water" "https://leetcode.com/problems/container-with-most-water/description/"
# Run from repo root. Creates new/YYYYMMDD/SLUG.md and src/SLUG/main.go (Monday = first day of current week).
# Idempotent: skips if files already exist.

set -e

TITLE="${1:?Usage: $0 \"<title>\" \"<url>\"}"
URL="${2:?Usage: $0 \"<title>\" \"<url>\"}"

# Monday of current week (YYYYMMDD), portable (Linux + macOS)
MONDAY=$(python3 -c "from datetime import datetime, timedelta; d=datetime.now(); m=d-timedelta(days=d.weekday()); print(m.strftime('%Y%m%d'))")

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

# Create note (idempotent)
mkdir -p "$NEW_DIR"
if [[ -f "$MD_FILE" ]]; then
  echo "Skipped (exists): $MD_FILE"
else
  cat > "$MD_FILE" << EOF
# ${TITLE}

## Link

[Leetcode](${URL})

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
  echo "Created: $MD_FILE"
fi

# Create src folder and main.go (idempotent)
if [[ -f "$MAIN_GO" ]]; then
  echo "Skipped (exists): $MAIN_GO"
else
  mkdir -p "$SRC_DIR"
  cat > "$MAIN_GO" << 'GOEOF'
package main

func main() {
}
GOEOF
  echo "Created: $MAIN_GO"
fi
