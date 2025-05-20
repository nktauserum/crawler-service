package task

import (
	"testing"

	"github.com/nktauserum/crawler-service/pkg/db"
)

func TestPendingTask(t *testing.T) {
	s := db.GetStorage()
	if s == nil {
		t.Fatal()
	}

	_, err := db.NewTask(s, "https://github.com/nktauserum/crawler-service")
	if err != nil {
		t.Fatal(err)
	}

	task, err := PendingTask(s)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(task)
}
