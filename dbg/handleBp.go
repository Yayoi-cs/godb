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

func (bp *TypeBp) EnableBp() error {
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

func (bp *TypeBp) DisableBp() error {
	int3InstrLittle := make([]byte, 8)
	_, err := syscall.PtracePeekData(bp.pid, bp.addr, int3InstrLittle)
	if err != nil {
		return err
	}
	int3Instr := binary.LittleEndian.Uint64(int3InstrLittle)
	origInstr := binary.LittleEndian.Uint64(bp.instr)
	newInstr := (int3Instr & ^uint64(0xff)) | (origInstr & 0xff)
	binInstr := make([]byte, 8)
	binary.LittleEndian.PutUint64(binInstr, newInstr)
	_, err = syscall.PtracePokeData(bp.pid, bp.addr, binInstr)
	if err != nil {
		return err
	}
	bp.isEnable = false
	return nil
}

func (dbger *TypeDbg) hitBp() error {
	rip, err := dbger.GetRip()
	if err != nil {
		return err
	}
	err = dbger.SetRip(rip - 1)
	if err != nil {
		return err
	}
	fmt.Printf("hit breakpoint at 0x%lx\n", rip-1)
	return nil
}

func newBp(bpAddr uintptr, pid int) (*TypeBp, error) {
	bp := &TypeBp{
		pid:   pid,
		addr:  bpAddr,
		instr: make([]byte, 8),
	}
	if err := bp.EnableBp(); err != nil {
		return nil, err
	}
	bp.isEnable = true
	return bp, nil
}
