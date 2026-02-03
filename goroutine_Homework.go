package main

import (
	"fmt"
	// "time"
)
func main()  {
	// go Go()
	// time.Sleep(2*time.Second)

	c:= make(chan string)

	go func (){
		i:=0
		for{
			select {
			case v,ok := <-c :
				if(ok){
					fmt.Println("iv",i,v)
					i++
				}
				if !ok {
					fmt.Println("channel closed")
					return
				}

			}
		}

	}()
	c<-"1"
	c<-"11"
	c<-"111"
	c<-"1111"
	close(c)

}
func Go(){
	fmt.Println("go go go !!!")
}