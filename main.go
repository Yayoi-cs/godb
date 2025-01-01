package main

import (
	"fmt"
	"godb/dbg"
)

func main() {
	fmt.Println("[*]main")
	dbger, err := dbg.Run("tmp", "a", "b")
	if err != nil {
		fmt.Println(err)
		return
	}
}
