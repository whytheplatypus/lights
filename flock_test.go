package lights_test

import (
	"os"
	"testing"
	"time"

	"github.com/whytheplatypus/lights"
)

func TestFlock(t *testing.T) {
	f, err := os.OpenFile("/run/lock/lights/razer.lights.lock", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		t.Fatal(err)
	}
	fd := int(f.Fd())
	l := &lights.Flock{fd}
	l.Lock()

	done := make(chan struct{})
	go func(done chan struct{}) {
		defer close(done)
		var l lights.Flock
		l.Lock()
		l.Unlock()
	}(done)
	l.Unlock()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("Failed to finish locking and unlocking in time")
	}
}

func TestFlock_HoldLock(t *testing.T) {
	f, err := os.OpenFile("/run/lock/lights/razer.lights.lock", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		t.Fatal(err)
	}
	fd := int(f.Fd())
	l := &lights.Flock{fd}
	l.Lock()
	defer l.Unlock()

	done := make(chan struct{})
	go func(done chan struct{}) {
		defer close(done)
		var l lights.Flock
		l.Lock()
		l.Unlock()
	}(done)
	select {
	case <-done:
		t.Fatal("Failed to hold the lock")
	case <-time.After(1 * time.Second):
	}
}
