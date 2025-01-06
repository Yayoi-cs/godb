package dbg

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func (dbger *TypeDbg) wait() (unix.WaitStatus, error) {
	var ws unix.WaitStatus
	_, err := unix.Wait4(0, &ws, unix.WALL, nil)
	if err != nil {
		return 0, err
	}
	if ws.Exited() {
		fmt.Println("target exited")
		return 0, fmt.Errorf("wait exited")
	}
	if ws.Stopped() {
		rip, err := dbger.GetRip()
		if err != nil {
			return 0, err
		}
		_, ok := dbger.bps[uintptr(rip)]
		if ok {
			fmt.Printf("[-]reach2breakpoint at %x\n", rip)
		} else {
			fmt.Printf("[-]Dbger Stopped at %x\n", rip)
		}
	}
	return ws, nil
}

func (dbger *TypeDbg) Wait() (unix.WaitStatus, error) {
	return dbger.wait()
}

func (dbger *TypeDbg) Continue() error {
	rip, err := dbger.GetRip()
	if err != nil {
		return err
	}
	bp, ok := dbger.bps[uintptr(rip)]
	if ok {
		if bp.isEnable {
			if err := bp.DisableBp(); err != nil {
				return err
			}
			if err := dbger.SetRip(rip - 1); err != nil {
				return err
			}
			if err := dbger.Step(); err != nil {
				return err
			}
			if _, err = dbger.wait(); err != nil {
				return err
			}
			if err := bp.EnableBp(); err != nil {
				return err
			}
		}
	}
	err = unix.PtraceCont(dbger.pid, 0)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) Step() error {
	err := unix.PtraceSingleStep(dbger.pid)
	if err != nil {
		fmt.Errorf("[-]failed to stepi: %w", err)
	}
	return nil
}
