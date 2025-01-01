package dbg

import (
	"encoding/binary"
	"fmt"
	"syscall"
)

func NewBp[T string | uintptr](bpAddr T, pid int) (*TypeBp, error) {
	switch v := any(bpAddr).(type) {
	case uintptr:
		nbp, err := newBp(v, pid)
		if err != nil {
			return nil, err
		}
		return nbp, nil
	case string:
		fmt.Println("to be continued")
	}
	return nil, nil
}

type TypeBp struct {
	pid    int
	addr   uintptr
	instr  []byte
	enable bool
}

func enable(bp *TypeBp) error {
	_, err := syscall.PtracePeekData(bp.pid, bp.addr, bp.instr)
	if err != nil {
		return err
	}
	origInstr := binary.LittleEndian.Uint64(bp.instr)
	int3Instr := (origInstr & ^uint64(0xff)) | 0xcc
	int3InstrLittle := make([]byte, 8)
	binary.LittleEndian.PutUint64(int3InstrLittle, int3Instr)
	_, err = syscall.PtracePokeData(bp.pid, bp.addr, int3InstrLittle)
	if err != nil {
		return err
	}

	return nil
}

func newBp(bpAddr uintptr, pid int) (*TypeBp, error) {
	bp := &TypeBp{
		pid:   pid,
		addr:  bpAddr,
		instr: make([]byte, 8),
	}
	if err := enable(bp); err != nil {
		return nil, err
	}
	bp.enable = true
	return bp, nil
}
