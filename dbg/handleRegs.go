package dbg

import (
	"errors"
	"syscall"
)

func GetRegs(pid int) (*syscall.PtraceRegs, error) {
	var ws syscall.WaitStatus
	_, err := syscall.Wait4(pid, &ws, syscall.WNOHANG, nil)
	if err != nil {
		return nil, err
	}
	if !ws.Stopped() && ws.Signal() != syscall.SIGSTOP {
		return nil, errors.New(plsStop)
	}

	regs := &syscall.PtraceRegs{}
	err = syscall.PtraceGetRegs(pid, regs)
	if err != nil {
		return nil, errors.New(stWrong)
	}

	return regs, nil
}

func GetR15(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R15, nil
}

func GetR14(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R14, nil
}

func GetR13(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R13, nil
}

func GetR12(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R12, nil
}

func GetRbp(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rbp, nil
}

func GetRbx(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rbx, nil
}

func GetR11(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R11, nil
}

func GetR10(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R10, nil
}

func GetR9(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R9, nil
}

func GetR8(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.R8, nil
}

func GetRax(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rax, nil
}

func GetRcx(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rcx, nil
}

func GetRdx(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rdx, nil
}

func GetRsi(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rsi, nil
}

func GetRdi(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rdi, nil
}

func GetOrig_rax(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Orig_rax, nil
}

func GetRip(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rip, nil
}

func GetCs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Cs, nil
}

func GetEflags(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Eflags, nil
}

func GetRsp(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Rsp, nil
}

func GetSs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Ss, nil
}

func GetFs_base(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Fs_base, nil
}

func GetGs_base(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Gs_base, nil
}

func GetDs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Ds, nil
}

func GetEs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Es, nil
}

func GetFs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Fs, nil
}

func GetGs(pid int) (uint64, error) {
	regs, err := GetRegs(pid)
	if err != nil {
		return 0, err
	}
	return regs.Gs, nil
}

func SetRegs(pid int, regs *syscall.PtraceRegs) error {
	err := syscall.PtraceSetRegs(pid, regs)
	if err != nil {
		return errors.New(stWrong)
	}
	return nil
}

func SetR15(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R15 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR14(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R14 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR13(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R13 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR12(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R12 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRbp(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rbp = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRbx(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rbx = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR11(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R11 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR10(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R10 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR9(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R9 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetR8(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.R8 = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRax(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rax = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRcx(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rcx = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRdx(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rdx = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRsi(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rsi = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRdi(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rdi = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetOrig_rax(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Orig_rax = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRip(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rip = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetCs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Cs = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetEflags(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Eflags = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetRsp(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Rsp = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetSs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Ss = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetFs_base(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Fs_base = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetGs_base(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Gs_base = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetDs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Ds = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetEs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Es = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetFs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Fs = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}

func SetGs(pid int, val uint64) error {
	regs, err := GetRegs(pid)
	if err != nil {
		return err
	}
	regs.Gs = val
	err = SetRegs(pid, regs)
	if err != nil {
		return err
	}
	return nil
}
