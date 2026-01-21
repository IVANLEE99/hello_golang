package main

import (
	"fmt";
)

type  test struct{

}
type person struct {
	Name string
	Age int
}

type person2 struct {
	Name string
	Age int
	Contact struct {
		Phone,City string
	}
}
func main()  {
	a:=test{}
	fmt.Println(a)

	b:=person{}
	b.Name = "Name"
	b.Age = 17
	fmt.Println(b)

	c:=person{Name:"sdf",Age: 77}
	fmt.Println(c)
	A(c)
	fmt.Println(c)
	B(&c)
	fmt.Println(c)

	d:=&person{Name:"dfdfdfdf",Age:999}
	d.Name = "ok"
	fmt.Println("d:",d)
	C(d)
	D(d)
	fmt.Println("d:",d)

	e:= struct {
		Name string
		Age int
	} {
		Name:"youngs",
		Age:18,
	}
	fmt.Println("e:",e)

	f:= person2{
		Name:"person2",
		Age:98,
	}
	f.Contact.Phone="121212"
	f.Contact.City="Guangzhou"
	fmt.Println("f:",f)


	type human struct {
		Sex int
	}

	type teacher struct {
		human
		Name string
		Age int
	}

	type student struct {
		human
		Name string
		Age int
	}

	g:=teacher{
		Name:"teacher",
		Age:77,
		human:human{
			Sex:1,
		},
	}

	h:=student{
		Name:"student",
		Age:70,
		human:human{
			Sex:0,
		},
	}
	g.Name="gg"
	h.Age = 88
	g.Sex=6666

	fmt.Println(g,h)











	
}
func A(per person){
	per.Age = 99
	fmt.Println(per)
}
func B(per *person){
	per.Age = 99
	fmt.Println(per)
}
func C(per *person){
	per.Age = 9
	fmt.Println(per)
}
func D(per *person){
	per.Age = 99000
	fmt.Println(per)
}