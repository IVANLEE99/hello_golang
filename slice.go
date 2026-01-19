package main

import (
	"fmt"
)

func main()  {
	a :=[10]int {}
	fmt.Println(a)
	s1 := a[5:10]
	fmt.Println(s1)
	s2 := make([]int,3,10)
	fmt.Println(len(s2),cap(s2))
	s3 := make([]int,3,6)
	fmt.Printf("%p\n",s3)
	s3 = append(s3, 1, 2, 3)
	fmt.Printf("%v %p\n",s3,s3)
	b := []int{1,2,3,4,5}
	s4:=b[2:5]
	s5:=b[1:3]
	fmt.Println(s4,s5)
	s5 = append(s2,1,2,1,1,1,1,1,1,1,1,1,1,1)
	s4[0]=9
	fmt.Println(s4,s5)
	s6:=[]int{1,2,3,4,5,6}
	s7:=[]int{7,8,9,10,1,1,1,1,1,1,1,1}
	// copy(s7,s6)
	// fmt.Println(s7)
	copy(s7[2:4],s6[1:3])
	fmt.Println(s7)

	c:= [...]int{1,2,3,4,5,6,7,8,9,0}
	d:=len(c)
	s8:= c[0:d]
	fmt.Println(s8)

	s9:=[]int{1,2,3,4,5,6,7,8}
	s10:=s9
	s11:=s9[:]
	fmt.Println(s10,s11)
}
