package lights

import (
	"log"
	"os"
	"syscall"
)

const lockfile = "/run/lock/lights/razer.lights.lock"

var fd int

func init() {
	f, err := os.OpenFile(lockfile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(0600))
	if err != nil {
		log.Println(err)
		return
	}
	fd = int(f.Fd())
}

type Flock struct{ Fd int }

func (f *Flock) Lock() {
	if f.Fd == 0 {
		f = &Flock{fd}
	}
	if err := syscall.Flock(f.Fd, syscall.LOCK_EX); err != nil {
		log.Println(err)
	}
}

func (f *Flock) Unlock() {
	if f.Fd == 0 {
		f = &Flock{fd}
	}
	if err := syscall.Flock(f.Fd, syscall.LOCK_UN); err != nil {
		log.Println(err)
	}
}
