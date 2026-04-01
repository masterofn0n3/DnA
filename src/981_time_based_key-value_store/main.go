package main

type Entry struct {
	value      string
	timestampt int
}
type TimeMap struct {
	timeMap map[string][]*Entry
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: make(map[string][]*Entry),
	}
}

// TimeMap timeMap = new TimeMap();
// timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
// timeMap.get("foo", 1);         // return "bar"
// timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
// timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
// timeMap.get("foo", 4);         // return "bar2"
// timeMap.get("foo", 5);         // return "bar2"
func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMap[key] = append(this.timeMap[key], &Entry{value, timestamp})
}

// 6
// [1, 4, 5, 8, 9]
func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.timeMap[key]
	if !ok {
		return ""
	}

	left, right := 0, len(entries)-1
	result := ""
	for left <= right {
		mid := (left + right) / 2
		if entries[mid].timestampt > timestamp {
			right = mid - 1
		} else {
			result = entries[mid].value
			left = mid + 1
		}
	}
	return result
}
