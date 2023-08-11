package main

import "fmt"

func main() {
	s := "hello,世界!"
	var a []byte
	a = []byte(s)
	var b string
	b = string(a)
	var c []rune
	c = []rune(s)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(a, b, c)
}
