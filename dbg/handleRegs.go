package dbg

import (
	"errors"
	"golang.org/x/sys/unix"
)

func (dbger *TypeDbg) GetRegs() (*unix.PtraceRegs, error) {
	pid := dbger.pid
	var ws unix.WaitStatus
	_, err := unix.Wait4(pid, &ws, unix.WNOHANG, nil)
	if err != nil {
		return nil, err
	}
	if !ws.Stopped() && ws.Signal() != unix.SIGSTOP {
		return nil, errors.New(plsStop)
	}

	regs := &unix.PtraceRegs{}
	err = unix.PtraceGetRegs(pid, regs)
	if err != nil {
		return nil, errors.New(stWrong)
	}

	return regs, nil
}

func (dbger *TypeDbg) GetR15() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R15, nil
}

func (dbger *TypeDbg) GetR14() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R14, nil
}

func (dbger *TypeDbg) GetR13() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R13, nil
}

func (dbger *TypeDbg) GetR12() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R12, nil
}

func (dbger *TypeDbg) GetRbp() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rbp, nil
}

func (dbger *TypeDbg) GetRbx() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rbx, nil
}

func (dbger *TypeDbg) GetR11() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R11, nil
}

func (dbger *TypeDbg) GetR10() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R10, nil
}

func (dbger *TypeDbg) GetR9() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R9, nil
}

func (dbger *TypeDbg) GetR8() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.R8, nil
}

func (dbger *TypeDbg) GetRax() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rax, nil
}

func (dbger *TypeDbg) GetRcx() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rcx, nil
}

func (dbger *TypeDbg) GetRdx() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rdx, nil
}

func (dbger *TypeDbg) GetRsi() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rsi, nil
}

func (dbger *TypeDbg) GetRdi() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rdi, nil
}

func (dbger *TypeDbg) GetOrig_rax() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Orig_rax, nil
}

func (dbger *TypeDbg) GetRip() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rip, nil
}

func (dbger *TypeDbg) GetCs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Cs, nil
}

func (dbger *TypeDbg) GetEflags() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Eflags, nil
}

func (dbger *TypeDbg) GetRsp() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Rsp, nil
}

func (dbger *TypeDbg) GetSs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Ss, nil
}

func (dbger *TypeDbg) GetFs_base() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Fs_base, nil
}

func (dbger *TypeDbg) GetGs_base() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Gs_base, nil
}

func (dbger *TypeDbg) GetDs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Ds, nil
}

func (dbger *TypeDbg) GetEs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Es, nil
}

func (dbger *TypeDbg) GetFs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Fs, nil
}

func (dbger *TypeDbg) GetGs() (uint64, error) {
	regs, err := dbger.GetRegs()
	if err != nil {
		return 0, err
	}
	return regs.Gs, nil
}

func (dbger *TypeDbg) SetRegs(regs *unix.PtraceRegs) error {
	pid := dbger.pid
	err := unix.PtraceSetRegs(pid, regs)
	if err != nil {
		return errors.New(stWrong)
	}

	return nil
}

func (dbger *TypeDbg) SetR15(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R15 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR14(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R14 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR13(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R13 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR12(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R12 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRbp(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rbp = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRbx(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rbx = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR11(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R11 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR10(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R10 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR9(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R9 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetR8(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.R8 = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRax(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rax = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRcx(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rcx = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRdx(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rdx = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRsi(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rsi = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRdi(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rdi = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetOrig_rax(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Orig_rax = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRip(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rip = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetCs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Cs = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetEflags(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Eflags = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetRsp(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Rsp = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetSs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Ss = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetFs_base(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Fs_base = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetGs_base(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Gs_base = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetDs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Ds = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetEs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Es = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetFs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Fs = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) SetGs(val uint64) error {
	regs, err := dbger.GetRegs()
	if err != nil {
		return err
	}
	regs.Gs = val
	err = dbger.SetRegs(regs)
	if err != nil {
		return err
	}
	return nil
}
