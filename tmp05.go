package main

import (
	"fmt"
	// "strconv"
)

func main()  {
	const (
		KB float64 = (1<<(iota*10))
		MB
		GB
		TB
		PB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	a:=1
	a++
	var p *int = &a//指针
	fmt.Println(p)
	fmt.Println(*p)//取值

	b:=10
	if b:=1;b>0{
		fmt.Println(b)
	}
	fmt.Println(b)


}