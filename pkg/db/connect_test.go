package db

import "testing"

func TestConnectToDB(t *testing.T) {
	s := GetStorage()
	if s == nil {
		t.Fail()
	}
}
