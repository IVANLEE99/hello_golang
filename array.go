package main

import (
	"fmt"
)

func main()  {
	var a[2] int
	var b[2] int
	b=a
	fmt.Println(b)
	
	c:=[20]int {10:1}
	fmt.Println(c)

	d:=[...]int {1,2,3,4,5}
	fmt.Println(d)
	
	e:=[...]int {99:1}
	fmt.Println(e)

	f:=[10]int{}
	f[2]=2
	fmt.Println(f)
	p:=new([10]int)
	p[1]=2
	fmt.Println(p)


}
