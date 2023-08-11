package main

import (
	"fmt"
)

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type MyInt int

func main() {
	var a MyInt = 10
	var b MyInt = 10

	c := a + b
	d := a * b

	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)

	var s Map = make(map[string]string)
	s["1"] = "11"
	s["2"] = "22"
	s["3"] = "33"
	s.Print()
}
