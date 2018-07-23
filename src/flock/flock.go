package flock

import "errors"
import "os"
import "time"

func Lock(spath string) (bool, error) {
	return LockWithTimeout(spath, time.Second*0)
}

func Unlock(spath string) error {
	err := os.Remove(spath)
	return err
}

func LockWithTimeout(spath string, timeout time.Duration) (bool, error) {
	start := time.Now()
	var t time.Time
	var err error
	var f *os.File
	for {
		f, err = os.OpenFile(spath, os.O_CREATE|os.O_EXCL, 0755)
		if err != nil {
			if !os.IsExist(err) {
				return false, err
			}
			t = time.Now()
			elapsed := t.Sub(start)
			if uint64(timeout) != uint64(time.Second*0) && elapsed > timeout {
				return false, errors.New("Lock Timeout")
			}
			time.Sleep(time.Microsecond)
			continue
		}
		break
	}
	err = f.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
