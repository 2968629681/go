package main

import (
	"fmt"
)

//defer func () {
//	println("defer inner")
//	recover()
//}()
//
//func except() {
//	recover()
//}
//
//func test() {
//	defer except()
//	panic("test panic")
//}
//
//test()

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//多个panic抛出时，只有最后一个能被捕获
	defer func() {
		panic("first")
	}()
	defer func() {
		panic("second")
	}()

	panic("main body")
}
