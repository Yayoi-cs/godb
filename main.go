package main

import (
	"fmt"
	"godb/dbg"
)

func main() {
	fmt.Println("[*]main")
	bin := "/home/tsuneki/dc/ctf/myctf/dbgerTest/dbgTest"
	dbger, err := dbg.Run(bin, true, "Alpaca")
	if err != nil {
		fmt.Println(err)
		return
	}
	dbger.Wait()
	fmt.Println(dbger.GetRegs())
	_, err = dbger.Break(0x1169)
	if err != nil {
		fmt.Println(err)
	}
	dbger.Continue()
	dbger.Wait()
	fmt.Println(dbger.GetRegs())
	dbger.Continue()
	return
}
