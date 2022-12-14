package longestsubstring

import "strings"

func lengthOfLongestSubstring(s string) int {
	var substr string
	var length int
	var maxLength int

	var i int
	l := len(s)

	for i < l {
		if si := strings.Index(substr, string(s[i])); si != -1 {
			substr = substr[si+1:]

			if maxLength < length {
				maxLength = length
			}
			length = len(substr)

		} else {
			substr += string(string(s[i]))
			length += 1
			i++
		}
	}

	if maxLength < length {
		maxLength = length
	}

	return maxLength
}
