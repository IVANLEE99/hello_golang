package main

import (
	"fmt";
	"sort"
)

func main()  {
	var m1 map[int]string
	m1 = map[int]string {}
	fmt.Println(m1)

	var m2 map[int]string = make(map[int]string)
	fmt.Println(m2)

	m3:=map[int]string {}
	fmt.Println(m3)

	m4:=make(map[int]string)
	m4[1]="ok"
	fmt.Println(m4[1])
	delete(m4,1)
	a:=m4[1]
	fmt.Println(a)

	var m5 map[int]map[int]string
	m5 = make(map[int]map[int]string)
	m5[1] = make(map[int]string)
	m5[1][1]="okkk"
	b:=m5[1][1]
	fmt.Println(b)

	var m6 map[int]map[int]string
	m6=make(map[int]map[int]string)
	c,ok := m6[2][1]
	if !ok {
		m6[2] = make(map[int]string)
	}
	m6[2][1]="okkkkk"
	c,ok = m6[2][1]
	fmt.Println(c,ok)

	sm:=make([]map[int]string,5)
	for i,v := range sm {
		v = make(map[int]string,1)
		v[1]="OKKKKKK"
		sm[i] = make(map[int]string,1)
		sm[i][1]="OKKKKKKi"
		fmt.Println(v)
	}
	fmt.Println(sm)


	m7:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f"}
	sm7 := make([]int,len(m7))
	si:=0
	for k,_ := range m7 {
		sm7[si]=k
		si++
	}
	fmt.Println(sm7)
	sort.Ints(sm7)
	fmt.Println(sm7)


	m8:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f"}
	m9:=make(map[string]int,len(m8))
	for k,v := range m8{
		m9[v]=k
	}
	fmt.Println(m8)
	fmt.Println(m9)

}
