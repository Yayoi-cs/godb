package main

import (
	"context"
	"fmt"
	"godb/dbg"
	"sync"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ{}"
)

type typeRes struct {
	c      byte
	status int
}

func analyze(wg *sync.WaitGroup, retChan chan typeRes, ctx context.Context, id int, flag string) {
	defer wg.Done()
	bin := "~/dc/ctf/alpaca/r8/hidden"
	done := make(chan struct{})
	go func() {
		defer close(done)

	outerloop:
		for j := 0; j < 10; j++ {
			dbger, err := dbg.Run(bin, true, flag)
			if err != nil {
				fmt.Println(err)
				continue
			}
			dbger.Wait()
			_, err = dbger.Break(0x151c)
			if err != nil {
				fmt.Println(err)
				continue
			}
			dbger.Continue()
			dbger.Wait()
			rdi, err := dbger.GetRdi()
			if err != nil {
				fmt.Printf("[%x]%v\n", id, err)
				continue
			}
			rsi, err := dbger.GetRsi()
			if err != nil {
				fmt.Printf("[%x]%v\n", id, err)
				continue
			}
			//fmt.Printf("rdi: %x ,rsi: %x\n", rdi, rsi)
			var i int
			lflag := len(flag)
			for i = 0; i < lflag+2; i++ {
				rdiVal, err := dbger.GetMemory1(uintptr(rdi + uint64(i)))
				if err != nil {
					fmt.Printf("[%x]%v\n", id, err)
					continue outerloop
				}
				rsiVal, err := dbger.GetMemory1(uintptr(rsi + uint64(i)))
				if err != nil {
					fmt.Printf("[%x]%v\n", id, err)
					continue outerloop
				}
				//fmt.Printf("%x:[%x]%x ,[%x]%x\n", id, rdi+uint64(i), rdiVal, rsi+uint64(i), rsiVal)
				if rdiVal != rsiVal || rdiVal == 0 || rsiVal == 0 {
					break
				}
			}
			retChan <- typeRes{flag[lflag-1], i}
			dbger.Continue()
			fmt.Printf("[%x]End analysis\n", id)
			break
		}
	}()
	select {
	case <-ctx.Done():
	case <-done:
	}
}

func main() {
	fmt.Println("[*]main")
	flag := "Alpaca{"
	for {
		res := make(chan typeRes, len(charset))
		var wg sync.WaitGroup
		for i, c := range charset {
			fmt.Println(flag + string(c))
			time.Sleep(time.Millisecond * 10)
			f := flag + string(c)
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()
			wg.Add(1)
			go analyze(&wg, res, ctx, i, f)
		}
		wg.Wait()
		close(res)
		var max typeRes = typeRes{
			c:      byte(0),
			status: 0,
		}
		for r := range res {
			if r.status > max.status {
				max.c = r.c
				max.status = r.status
			}
		}
		flag = flag + string(max.c)
		if string(max.c) == "}" {
			break
		}
	}
	fmt.Println("[*]Shutdown godb")
	return
}
