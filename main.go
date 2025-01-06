package main

import (
	"fmt"
	"godb/dbg"
	"sync"
	"time"
)

func analyze(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	bin := "/home/tsuneki/dc/ctf/myctf/dbgerTest/dbgTest"
	dbger, err := dbg.Run(bin, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbger.Wait()
	//fmt.Println(dbger.GetRegs())
	_, err = dbger.Break(0x1169)
	if err != nil {
		fmt.Println(err)
	}
	dbger.Continue()
	dbger.Wait()
	fmt.Println(dbger.GetRegs())
	dbger.Continue()
	fmt.Printf("[%x]wait4child\n", id)
	dbger.Wait()
	fmt.Printf("[%x]End analysis\n", id)
}

func main() {
	fmt.Println("[*]main")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		time.Sleep(20)
		wg.Add(1)
		go analyze(&wg, i)
	}
	wg.Wait()
	fmt.Println("[*]Shutdown godb")
	return
}
