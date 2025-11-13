package main

import (
	"fmt"
	basicreverse "task-4/basicReverse"
	"task-4/smartreverse"
)

func main() {
	s := "abc defg higk lmnop"
	res1 := basicreverse.BasicReverse(s)
	res2 := smartreverse.SmartReverse(s)
	fmt.Println(res1)
	fmt.Println(res2)
}


