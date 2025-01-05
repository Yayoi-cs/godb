package dbg

import (
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

func (dbger *TypeDbg) GetMemory(n uint, addr uintptr) ([]byte, error) {
	mem := make([]byte, n)
	count, err := unix.PtracePeekData(dbger.pid, addr, mem)
	if err != nil {
		return nil, err
	}
	if uint(count) != n {
		return nil, errors.New(stWrong)
	}
	return mem, nil
}

func (dbger *TypeDbg) GetMemory1(addr uintptr) (uint8, error) {
	mem, err := dbger.GetMemory(1, addr)
	if err != nil {
		return 0, err
	}
	return uint8(mem[0]), nil
}

func (dbger *TypeDbg) GetMemory2(addr uintptr) (uint16, error) {
	mem, err := dbger.GetMemory(2, addr)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(mem), nil
}

func (dbger *TypeDbg) GetMemory4(addr uintptr) (uint32, error) {
	mem, err := dbger.GetMemory(4, addr)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(mem), nil
}

func (dbger *TypeDbg) GetMemory8(addr uintptr) (uint64, error) {
	mem, err := dbger.GetMemory(8, addr)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(mem), nil
}

func (dbger *TypeDbg) SetMemory(data []byte, addr uintptr) error {
	count, err := unix.PtracePokeData(dbger.pid, addr, data)
	if err != nil {
		return err
	}
	if count != len(data) {
		return errors.New(stWrong)
	}
	return nil
}

func (dbger *TypeDbg) SetMemory1(data uint8, addr uintptr) error {
	bData := []byte{data}
	return dbger.SetMemory(bData, addr)
}

func (dbger *TypeDbg) SetMemory2(data uint16, addr uintptr) error {
	bData := make([]byte, 2)
	binary.LittleEndian.PutUint16(bData, data)
	return dbger.SetMemory(bData, addr)
}

func (dbger *TypeDbg) SetMemory4(data uint32, addr uintptr) error {
	bData := make([]byte, 4)
	binary.LittleEndian.PutUint32(bData, data)
	return dbger.SetMemory(bData, addr)
}

func (dbger *TypeDbg) SetMemory8(data uint64, addr uintptr) error {
	bData := make([]byte, 8)
	binary.LittleEndian.PutUint64(bData, data)
	return dbger.SetMemory(bData, addr)
}
