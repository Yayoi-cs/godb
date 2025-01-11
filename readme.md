# godb
the fastest debugger for brute-force
## benchmarks
x15 faster than gdb script.
* alpacaHack round8
  * launch 6,890 process and over 361,725 memory access in 85s
  * example in main.go
## example
### run
```go
dbger, err := dbg.Run("/path/to/binary", true, "command","line","arguments")
```
### set breakpoint
```go
_, err := dbger.Break(0x151c) //relative address from base address
```
### get/set registers
```go
//get rdi -> uint64
rdi, err := dbger.GetRdi()
//get rsi -> uint64
rsi, err := dbger.GetRsi()
fmt.Printf("rdi: %x ,rsi: %x\n", rdi, rsi)
//set rip
if err := dbger.SetRip(rip - 1); err != nil {
    fmt.Println(err)
}
```
### get/set memory
```go
//get 1 byte -> uint8
rdiVal, err := dbger.GetMemory1(uintptr(rdi))
//get 2 bytes -> uint16
rdiVal, err := dbger.GetMemory2(uintptr(rdi))
//get 4 bytes -> uint32
rdiVal, err := dbger.GetMemory4(uintptr(rdi))
//get 8 bytes -> uint64
rdiVal, err := dbger.GetMemory8(uintptr(rdi))
```
### continue,step
```go
dbger.Continue()
```
### wait
```go
dbger.Wait()
```
### parallel debug
> godb is so fast. if launch many process in no wait, many error will be caused.
> so, I recommend to add try like method in analysis function with timeout.
```go
const (
	charset = "abcdefghijklmnopqrstuvwxyz0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ{}"
)
func analyze(wg *sync.WaitGroup, retChan chan int, ctx context.Context, argv string) {
	defer wg.Done()
	bin := "/path/to/your/binary"
	done := make(chan struct{})
	go func() {
		defer close(done)
		//try to start debugging up to 10 times.
		for j := 0; j < 10; j++ {
			dbger, err := dbg.Run(bin, true, argv)
			if err != nil { continue }
			if _,err := dbger.Wait();err != nil { continue }
			retChan <- i
			break
		}
	}()
	select {
	case <-ctx.Done():
	case <-done:
	}
}
func main() {
	res := make(chan int, len(charset))
	var wg sync.WaitGroup
	for i, c := range charset {
		time.Sleep(time.Millisecond * 10)
    	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		wg.Add(1)
		go analyze(&wg, res, ctx, string(c))
	}
	wg.Wait()
	close(res)
    return
}
```
### AlpacaHack Round8 hidden
It's possible to solve by brute force
Solver in main.go in master branch
[main.go:godb-master](https://github.com/Yayoi-cs/godb/blob/master/main.go)
## functions
```go
func Run(bin string, pie bool, args ...string) (*TypeDbg, error)
func (dbger *TypeDbg) LoadBase() error 

func (dbger *TypeDbg) Wait() (unix.WaitStatus, error)
func (dbger *TypeDbg) Continue() error
func (dbger *TypeDbg) Step() error

func (dbger *TypeDbg) GetRegs() (*unix.PtraceRegs, error)
func (dbger *TypeDbg) GetR15() (uint64, error)
func (dbger *TypeDbg) GetR14() (uint64, error)
func (dbger *TypeDbg) GetR13() (uint64, error)
func (dbger *TypeDbg) GetR12() (uint64, error)
func (dbger *TypeDbg) GetRbp() (uint64, error)
func (dbger *TypeDbg) GetRbx() (uint64, error)
func (dbger *TypeDbg) GetR11() (uint64, error)
func (dbger *TypeDbg) GetR10() (uint64, error)
func (dbger *TypeDbg) GetR9() (uint64, error)
func (dbger *TypeDbg) GetR8() (uint64, error)
func (dbger *TypeDbg) GetRax() (uint64, error)
func (dbger *TypeDbg) GetRcx() (uint64, error)
func (dbger *TypeDbg) GetRdx() (uint64, error)
func (dbger *TypeDbg) GetRsi() (uint64, error)
func (dbger *TypeDbg) GetRdi() (uint64, error)
func (dbger *TypeDbg) GetOrig_rax() (uint64, error)
func (dbger *TypeDbg) GetRip() (uint64, error)
func (dbger *TypeDbg) GetCs() (uint64, error)
func (dbger *TypeDbg) GetEflags() (uint64, error)
func (dbger *TypeDbg) GetRsp() (uint64, error)
func (dbger *TypeDbg) GetSs() (uint64, error)
func (dbger *TypeDbg) GetFs_base() (uint64, error)
func (dbger *TypeDbg) GetGs_base() (uint64, error)
func (dbger *TypeDbg) GetDs() (uint64, error)
func (dbger *TypeDbg) GetEs() (uint64, error)
func (dbger *TypeDbg) GetFs() (uint64, error)
func (dbger *TypeDbg) GetGs() (uint64, error)

func (dbger *TypeDbg) SetRegs(regs *unix.PtraceRegs) error
func (dbger *TypeDbg) SetR15(val uint64) error
func (dbger *TypeDbg) SetR14(val uint64) error
func (dbger *TypeDbg) SetR13(val uint64) error
func (dbger *TypeDbg) SetR12(val uint64) error
func (dbger *TypeDbg) SetRbp(val uint64) error
func (dbger *TypeDbg) SetRbx(val uint64) error
func (dbger *TypeDbg) SetR11(val uint64) error
func (dbger *TypeDbg) SetR10(val uint64) error
func (dbger *TypeDbg) SetR9(val uint64) error
func (dbger *TypeDbg) SetR8(val uint64) error
func (dbger *TypeDbg) SetRax(val uint64) error
func (dbger *TypeDbg) SetRcx(val uint64) error
func (dbger *TypeDbg) SetRdx(val uint64) error
func (dbger *TypeDbg) SetRsi(val uint64) error
func (dbger *TypeDbg) SetRdi(val uint64) error
func (dbger *TypeDbg) SetOrig_rax(val uint64) error
func (dbger *TypeDbg) SetRip(val uint64) error
func (dbger *TypeDbg) SetCs(val uint64) error
func (dbger *TypeDbg) SetEflags(val uint64) error
func (dbger *TypeDbg) SetRsp(val uint64) error
func (dbger *TypeDbg) SetSs(val uint64) error
func (dbger *TypeDbg) SetFs_base(val uint64) error
func (dbger *TypeDbg) SetGs_base(val uint64) error
func (dbger *TypeDbg) SetDs(val uint64) error
func (dbger *TypeDbg) SetEs(val uint64) error
func (dbger *TypeDbg) SetFs(val uint64) error
func (dbger *TypeDbg) SetGs(val uint64) error

func (dbger *TypeDbg) Break(bpAddr interface{}) (*TypeBp, error)
func (bp *TypeBp) EnableBp() error
func (bp *TypeBp) DisableBp() error

func (dbger *TypeDbg) GetMemory(n uint, addr uintptr) ([]byte, error)
func (dbger *TypeDbg) GetMemory1(addr uintptr) (uint8, error)
func (dbger *TypeDbg) GetMemory2(addr uintptr) (uint16, error)
func (dbger *TypeDbg) GetMemory4(addr uintptr) (uint32, error)
func (dbger *TypeDbg) GetMemory8(addr uintptr) (uint64, error)

func (dbger *TypeDbg) SetMemory(data []byte, addr uintptr) error
func (dbger *TypeDbg) SetMemory1(data uint8, addr uintptr) error
func (dbger *TypeDbg) SetMemory2(data uint16, addr uintptr) error
func (dbger *TypeDbg) SetMemory4(data uint32, addr uintptr) error
func (dbger *TypeDbg) SetMemory8(data uint64, addr uintptr) error

func (dbger *TypeDbg) SendLine(payload []byte) error
func (dbger *TypeDbg) Send(payload []byte) error
func (dbger *TypeDbg) Recv() ([]byte, error)
```
