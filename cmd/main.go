package main

import (
	"fmt"
	"postBoy/internal/search"
	"postBoy/internal/memory"
)

func main() {
	fmt.Println("数据是",search.A)
	fmt.Println("数组是",memory.Arr)
	fmt.Println("切片是",memory.Num)
	memory.Setnum()
	fmt.Println("切片是",memory.Setnum())
	fmt.Print("你好")
}
