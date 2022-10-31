package medium

import (
	"testing"
)

func TestOutersections(t *testing.T) {
	A := []string{"1", "2", "3", "4", "5", "6", "7"}
	B := []string{"5", "6", "7", "8", "9", "10"}
	expected := []string{"1", "2", "3", "4", "10", "8", "9"}

	r := Outersection(A, B)
	for i, e := range expected {
		if e != r[i] {
			t.Fail()
		}
	}
}
