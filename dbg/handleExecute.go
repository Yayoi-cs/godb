package dbg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type TypeDbg struct {
	pid   int
	path  string
	bases TypeAddr
	bps   map[uintptr]*TypeBp
}

type TypeAddr struct {
	bin  uintptr
	stk  uintptr
	heap uintptr
	libc uintptr
	ld   uintptr
}

func (dbger *TypeDbg) LoadBase() error {
	mapsFile := fmt.Sprintf("/proc/%d/maps", dbger.pid)
	fd, err := os.Open(mapsFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}
		binname := fields[5]
		if strings.HasPrefix(binname, dbger.path) && dbger.bases.bin != 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.bin = uintptr(addr)
		} else if strings.HasPrefix(binname, "[heap]") && dbger.bases.heap != 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.heap = uintptr(addr)
		} else if strings.HasPrefix(binname, "[stack]") && dbger.bases.stk != 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.stk = uintptr(addr)
		} else if strings.Contains(binname, "libc") && dbger.bases.libc != 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.libc = uintptr(addr)
		} else if strings.Contains(binname, "ld") && dbger.bases.ld != 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.ld = uintptr(addr)
		}
	}
	return nil
}

func Run(bin string, args ...string) (*TypeDbg, error) {
	absPath, err := filepath.Abs(bin)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(absPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("failed to start %v: %v", bin, err)
	}

	dbger := &TypeDbg{
		pid:  cmd.Process.Pid,
		path: absPath,
		bases: TypeAddr{
			bin:  0,
			stk:  0,
			heap: 0,
			libc: 0,
			ld:   0,
		},
		bps: make(map[uintptr]*TypeBp),
	}
	err = dbger.LoadBase()
	if err != nil {
		return nil, err
	}

	return dbger, nil
}

func Start() {

}
