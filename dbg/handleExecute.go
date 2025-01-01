package dbg

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

type TypeDbg struct {
	pid int
	bps map[uint64]*TypeBp
}

func Run(bin string, args ...string) (*TypeDbg, error) {
	cmd := exec.Command(bin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("failed to start %v: %v", bin, err)
	}

	dbger := &TypeDbg{
		pid: cmd.Process.Pid,
		bps: make(map[uint64]*TypeBp),
	}

	return dbger, nil
}

func Start() {

}
