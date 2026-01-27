package main

import (
	"fmt"
	"reflect"
)
type User struct {
	Id int
	Name string
	Age int
}
func (u User) Hello(){
	fmt.Println("hello world")
}
func main()  {
	u:= User{1,"ok",12}
	// u.Hello()
	Info(u)
	Set(&u)
	fmt.Println(u)

	x:=123
	xv:=reflect.ValueOf(&x)
	xv.Elem().SetInt(999)
	fmt.Println(x)
}

func Info (o interface{}){
	t:=reflect.TypeOf(o)
	fmt.Println("type",t.Name())

	if k:=t.Kind();k != reflect.Struct {
		fmt.Println("XXX")
		return
	}

	v:=reflect.ValueOf(o)
	fmt.Println("Fields",v)

	for i:=0;i<t.NumField();i++ {
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n",f.Name,f.Type,val)
	}

	for i:=0;i<t.NumMethod();i++ {
		m:=t.Method(i)
		fmt.Printf("%6s: %v\n",m.Name,m.Type)
	}
}


func Set(o interface{}){
	v:=reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("XXX. Set")
		return
	}else{
		v=v.Elem()
	}

	if f:=v.FieldByName("Name");f.Kind() == reflect.String{
		f.SetString("BYe")
	}
}