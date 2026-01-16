package main

import (
	"fmt"
	"strconv"
)

func main()  {
LABEL:
	for i:=0; i<10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
a:=1
for {
	a++
	if a>3 {
		break
	}

}
fmt.Println(strconv.Itoa('-'))
fmt.Println(a)

for a<10 {
	a++
}
fmt.Println('-')
fmt.Println(a)

for i:=0; i<10; i++ {
	a++
}
fmt.Println(a)




}