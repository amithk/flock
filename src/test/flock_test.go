package flocktest

import "errors"
import "testing"
import "flock"

func TestBasicFlockFunctionality(t *testing.T) {
	lpath := "/Users/personal/abc.lock"
	acquired, err := flock.Lock(lpath)
	if err != nil {
		t.Errorf("Error %v during Lock", err)
	}
	if !acquired {
		t.Errorf("Unexpected Error %v", errors.New("Wrong value of acquired"))
	}
	err = flock.Unlock(lpath)
	if err != nil {
		t.Errorf("Error %v during Unlock", err)
	}
}

