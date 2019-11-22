package main

import "testing"

func TestRepeatedString1(t *testing.T) {
	s := "aba"
	n := int64(10)

	numOfAs := repeatedString(s, n)
	if numOfAs != 7 {
		t.Errorf("got %d instead of 7", numOfAs)
	}

}
