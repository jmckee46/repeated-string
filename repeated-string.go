package main

import (
	"strings"
)

func repeatedString(s string, n int64) int64 {
	sLength := int64(len(s))

	numOfAs := int64(strings.Count(s, "a"))

	completeRepeats, remainder := n/sLength, n%sLength

	completeRepeatAs := numOfAs * completeRepeats

	remainingAs := int64(strings.Count(s[:remainder], "a"))

	return completeRepeatAs + remainingAs
}

func main() {

}
