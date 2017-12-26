package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, err := Division(6, 2); i != 3 || err != nil {
		t.Error("Tests failed")
	} else {
		t.Log("Tests passed")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("Not passed")
}
