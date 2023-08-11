package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

func doinput(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func wrap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//函数后面的小括号表示调用这个函数
	println(sum(5, 2))

	doinput(func(x, y int) int {
		return x + y
	}, 1, 2)

	opFunc := wrap("addfvc ")
	re := opFunc(2, 3)

	fmt.Printf("%d\n", re)
}
