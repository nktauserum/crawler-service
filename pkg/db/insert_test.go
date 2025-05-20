package db

import "testing"

func TestTaskInsertion(t *testing.T) {
	s := GetStorage()
	if s == nil {
		t.Fatal()
	}

	uuid, err := NewTask(s, "https://github.com/nktauserum/crawler-service")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(uuid)
}
