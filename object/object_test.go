package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{"hello"}
	hello2 := &String{"hello"}
	diff1 := &String{"diff"}
	diff2 := &String{"diff"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
