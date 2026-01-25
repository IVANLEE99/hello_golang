package main

import (
	"fmt";
)

type A struct{
	Name string
}

type B struct{
	Name string
}

type TZ  int

func main()  {
	a := A{}
	a.Print()

	b:= B{}
	b.Print()

	var tz TZ
	tz.Print()
	(*TZ).Print(&tz)

	tz.Increase(100)
	fmt.Println(tz)




}
func (a A) Print (){
	fmt.Println("A")
}

func (b B) Print (){
	fmt.Println("B")
}
func (a *TZ) Print(){
	fmt.Println("TZ")
}
func (tz *TZ)Increase(num int){
	*tz+=TZ(num)
}