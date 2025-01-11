package dbg

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type TypeDbg struct {
	pid          int
	path         string
	bases        TypeAddr
	stdinReader  *io.PipeReader
	stdinWriter  *io.PipeWriter
	stdoutReader *io.PipeReader
	stdoutWriter *io.PipeWriter
	bps          map[uintptr]*TypeBp
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
	dbger.bases = TypeAddr{
		bin:  0,
		stk:  0,
		heap: 0,
		libc: 0,
		ld:   0,
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) <= 5 {
			continue
		}
		binName := fields[5]

		if binName == dbger.path && dbger.bases.bin == 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.bin = uintptr(addr)
		} else if strings.Contains(binName, "[heap]") && dbger.bases.heap == 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.heap = uintptr(addr)
		} else if strings.Contains(binName, "[stack]") && dbger.bases.stk == 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.stk = uintptr(addr)
		} else if strings.Contains(binName, "libc") && dbger.bases.libc == 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.libc = uintptr(addr)
		} else if strings.Contains(binName, "ld") && dbger.bases.ld == 0 {
			addrArea := fields[0]
			addrStr := strings.Split(addrArea, "-")[0]
			addr, err := strconv.ParseUint(addrStr, 16, 64)
			if err != nil {
				return err
			}
			dbger.bases.ld = uintptr(addr)
		}
	}
	fmt.Printf("[-]target loaded at: %x\n", dbger.bases.bin)
	return nil
}

func Run(bin string, pie bool, pipe bool, args ...string) (*TypeDbg, error) {
	absPath := bin
	if strings.HasPrefix(bin, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		absPath = filepath.Join(home, bin[1:])
	}
	absPath, err := filepath.Abs(absPath)
	if err != nil {
		return nil, err
	}
	dbger := &TypeDbg{
		pid:  -1,
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
	cmd := exec.Command(absPath, args...)
	if pipe {
		stdinR, stdinW := io.Pipe()
		stdoutR, stdoutW := io.Pipe()
		dbger.stdinReader = stdinR
		dbger.stdinWriter = stdinW
		dbger.stdoutReader = stdoutR
		dbger.stdoutWriter = stdoutW
		cmd.Stdin = stdinR
		cmd.Stdout = stdoutW
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		dbger.stdinWriter = nil
		dbger.stdinReader = nil
		dbger.stdoutReader = nil
		dbger.stdoutWriter = nil
	}

	cmd.SysProcAttr = &unix.SysProcAttr{
		Ptrace: true,
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("failed to start %v: %v", bin, err)
	}
	dbger.pid = cmd.Process.Pid
	if pie {
		err = dbger.LoadBase()
	}
	if err != nil {
		return nil, err
	}
	fmt.Printf("[*]dbger started at pid: %d\n", dbger.pid)
	return dbger, nil
}
