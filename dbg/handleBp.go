package dbg

import (
	"encoding/binary"
	"errors"
	"fmt"
	"syscall"
)

func (dbger *TypeDbg) Break(bpAddr interface{}) (*TypeBp, error) {
	pid := dbger.pid
	switch v := bpAddr.(type) {
	case uintptr:
		nbp, err := newBp(v, pid)
		if err != nil {
			return nil, err
		}
		return nbp, nil
	case string:
		fmt.Println("to be continued")
	default:
		return nil, errors.New(typeError)
	}
	return nil, nil
}

type TypeBp struct {
	pid      int
	addr     uintptr
	instr    []byte
	isEnable bool
}

func (bp *TypeBp) enable() error {
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
	if err := bp.enable(); err != nil {
		return nil, err
	}
	bp.isEnable = true
	return bp, nil
}
