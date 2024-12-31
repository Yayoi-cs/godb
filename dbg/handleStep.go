package dbg

import (
	"fmt"
	"syscall"
)

func C(pid int) error {
	return Continue(pid)
}

func Continue(pid int) error {
	err := syscall.PtraceCont(pid, 0)
	if err != nil {
		fmt.Errorf("[-]failed to continue")
	}
	return nil
}

func S(pid int) error {
	return Step(pid)
}

func Step(pid int) error {
	err := syscall.PtraceSingleStep(pid)
	if err != nil {
		fmt.Errorf("[-]failed to stepi")
	}
	return nil
}
