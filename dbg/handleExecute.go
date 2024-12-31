package dbg

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Run(bin string, args ...string) (pid int, err error) {
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

	return cmd.Process.Pid, nil
}

func Start() {

}
