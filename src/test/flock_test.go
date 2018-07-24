package flocktest

import "testing"
import "flock"

func TestBasicFlockFunctionality(t *testing.T) {
	lpath := "/Users/personal/abc.lock"
	err := flock.Lock(lpath)
	if err != nil {
		t.Errorf("Error %v during Lock", err)
	}
	err = flock.Unlock(lpath)
	if err != nil {
		t.Errorf("Error %v during Unlock", err)
	}
}
