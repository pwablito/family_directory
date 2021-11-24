package util

import "testing"

func TestHash(t *testing.T) {
	test1Hash := HashString("test1")
	test2Hash := HashString("test2")
	if test1Hash == test2Hash {
		t.Error("digests should not be equal")
	}
	if test1Hash != HashString("test1") {
		t.Error("digests should be equal")
	}
}
