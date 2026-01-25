package main

import (
	"fmt";
)

type A struct{
	B
}

type B struct{
	Name string
}

func main()  {
	a := A{B: B{Name:"B"}}
	fmt.Println(a.Name, a.B.Name)

}