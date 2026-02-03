package main

import (
	"fmt"
	// "time"
)
func main()  {
	// go Go()
	// time.Sleep(2*time.Second)

	c:= make(chan bool)
	go func(){
		fmt.Println("go go go !!!")
		c<-true
		close(c)
	}()
	// <-c
	for v:=range c{
		fmt.Println("v",v)
	}

	



}
func Go(){
	fmt.Println("go go go !!!")
}