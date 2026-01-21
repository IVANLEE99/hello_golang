package main

import (
	"fmt";
)

func main()  {
	var fs = [4]func(){}


	for i:=0;i<4;i++{
		defer fmt.Println("defer i",i)
		defer func (){
			fmt.Println("defer clouser i",i)
		}()
		fs[i] = func() {fmt.Println(" clouser i",i)}
	}
	for _,f :=range fs{
		f()
	}

}
//  clouser i 0
//  clouser i 1
//  clouser i 2
//  clouser i 3
// defer clouser i 3
// defer i 3
// defer clouser i 2
// defer i 2
// defer clouser i 1
// defer i 1
// defer clouser i 0
// defer i 0