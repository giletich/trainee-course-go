package main

import (
	"fmt"
	basicreverse "task-4/basicrev"
	"task-4/smartreverse"
)

func main() {
	s := "abc defg higk lmnop пупупу"
	res1 := basicreverse.BasicReverse(s)
	res2 := smartreverse.SmartReverse(s)
	fmt.Println(res1)
	fmt.Println(res2)
}


