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
