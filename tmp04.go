package main

import (
	"fmt"
	// "strconv"
)

func main()  {
	const (
		a=1
		b=iota
		c
		d=iota
	)
	const (
		e=iota//重置
		f
	)
	const g=iota
	fmt.Println(a,b,c,d,e,f,g)
}