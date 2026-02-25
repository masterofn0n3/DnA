package main

import "strconv"

func main() {

}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	delim := "#"
	result := ""
	for _, v := range strs {
		result += delim + strconv.Itoa(len(v)) + delim + v
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	pos := 0
	for pos < len(encoded) {
		if encoded[pos] == '#' {
			var num []byte
			for j := pos + 1; j < len(encoded); j++ {
				if encoded[j] == '#' {
					break
				}
				num = append(num, encoded[j])
			}
			nextwordPos := pos + len(num) + 2
			nextwordLen, _ := strconv.Atoi(string(num))

			result = append(result, encoded[nextwordPos:nextwordPos+nextwordLen])
			pos = nextwordPos + nextwordLen
		}

	}
	return result
}

// func (s *Solution) Encode(strs []string) string {
// 	delim := "#"
// 	result := ""
// 	for _, v := range strs {
// 		result += strconv.Itoa(len(v)) + delim + v
// 	}
// 	return result
// }

// func (s *Solution) Decode(encoded string) []string {
// 	var result []string
// 	pos := 0
// 	for pos < len(encoded) {
// 		j := pos
// 		for encoded[j] != '#' {
// 			j++
// 		}
// 		lengthStr := encoded[pos:j]
// 		length, _ := strconv.Atoi(lengthStr)

// 		start := j + 1
// 		end := start + length
// 		result = append(result, encoded[start:end])
// 		pos = end
// 	}
// 	return result
// }
