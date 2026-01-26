package main

import (
	"fmt";
)
type empty interface{

}
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnectr struct {
	name string
}

func (pc PhoneConnectr) Name() string{
	return pc.name
}

func (pc PhoneConnectr) Connect(){
	fmt.Println("Connected:",pc.name)
}


func main()  {
	a:=PhoneConnectr{"PhoneConnectr"}
	a.Connect()
	Disconnect(a)
}

func Disconnect (usb USB){
	if pc,ok:= usb.(PhoneConnectr); ok {
		fmt.Println("Disconnected...",pc.name)
	}
	switch v:=usb.(type){
	case PhoneConnectr:
		fmt.Println("Disconnected...",v.name)
	default:
		fmt.Println("unknown device...")
	}

	pc2:=PhoneConnectr{"PhoneConnectr2"}
	var c2 Connecter
	c2=Connecter(pc2)
	c2.Connect()
}