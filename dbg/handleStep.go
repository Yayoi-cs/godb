package dbg

import (
	"fmt"
	"syscall"
)

func (dbger *TypeDbg) Continue() error {
	err := syscall.PtraceCont(dbger.pid, 0)
	if err != nil {
		fmt.Errorf("[-]failed to continue")
	}
	return nil
}

func (dbger *TypeDbg) Step() error {
	err := syscall.PtraceSingleStep(dbger.pid)
	if err != nil {
		fmt.Errorf("[-]failed to stepi")
	}
	return nil
}
