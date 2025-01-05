package dbg

import (
	"fmt"
	"syscall"
)

func (dbger *TypeDbg) wait() (syscall.WaitStatus, error) {
	var ws syscall.WaitStatus
	_, err := syscall.Wait4(0, &ws, syscall.WALL, nil)
	if err != nil {
		return 0, err
	}
	if ws.Exited() {
		return 0, fmt.Errorf("wait exited")
	}
	return ws, nil
}

func (dbger *TypeDbg) Continue() error {
	rip, err := dbger.GetRip()
	if err != nil {
		return err
	}
	bp, ok := dbger.bps[uintptr(rip)]
	if ok {
		if bp.isEnable {
			err = bp.DisableBp()
			if err != nil {
				return err
			}
			err = dbger.Step()
			if err != nil {
				return err
			}

		}
	}
	err = syscall.PtraceCont(dbger.pid, 0)
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
