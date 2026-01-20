package main

import (
	"fmt";
)

func main()  {
	a,b,c:=A()
	fmt.Println(a,b,c)
	B(1,2,3,4,5,6,7,8,9,0)

	e,f:=1,2
	C(e,f)
	fmt.Println("ef:",e,f)

	g1:=[]int{1,2,3,4}
	D(g1)
	fmt.Println("g1:",g1)

	h:=1
	F(h)
	fmt.Println("h:",h)

	i:=1
	G(&i)
	fmt.Println("i:",i)

	j:=H
	j()

	k:= func(){
		fmt.Println("K func")
	}
	k()

	l:= closure(10)
	fmt.Println(l(1))
	fmt.Println(l(2))


	fmt.Println("001")
	defer fmt.Println("002")
	defer fmt.Println("003")


	for i:=0;i<3;i++{
		defer fmt.Println("00I",i)
	}

	for i:=0;i<3;i++{
		defer func (){
			fmt.Println("defer func",i)
		}()
	}

	AA()
	BB()
	CC()



}

func A() (a,b,c int){
	a,b,c=1,2,3
	return
}

func B(d ...int){
	fmt.Println(d)
}

func C(s ...int){
	s[0]=0
	s[1]=1
	fmt.Println("c:",s)
}

func D(s []int){
	s[0]=5
	s[1]=6
	s[2]=7
	s[3]=8
	fmt.Println("D:",s)
}

func F(a int){
	a=2
	fmt.Println("F:",a)
}

func G (a *int){
	*a = 2
	fmt.Println("F:",*a)
}

func H(){
	fmt.Println("func H")
}

func closure(x int) func(int) int{
	fmt.Printf("%p\n",&x)
	return func (y int) int {
		fmt.Printf("%p\n",&x)
		return x+y
	}
}


func AA(){
	fmt.Println("FuncAA")
}

func BB(){
	defer func(){
		if err := recover();err != nil {
			fmt.Println("Recover in BB:",err)
		}
	}()
	panic("Panic in BB")
	fmt.Println("FuncBB")
}

func CC(){
	fmt.Println("FuncCC")
}