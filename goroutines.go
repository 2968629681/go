package main

import "runtime"

//func main() {
//	c := make(chan struct{})
//	go func(i chan struct{}) {
//
//		sum := 0
//		for i := 0; i < 10000; i++ {
//			sum += i
//		}
//		println(sum)
//		//写通道
//		c <- struct{}{}
//	}(c)
//
//	println("NumGoroutine=", runtime.NumGoroutine())
//	//读通道c，通过通道进行同步等待
//	<-c
//}

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		//写通道		匿名结构体赋值，前面的括号输入结构体的内如，后一个赋值  不能省略！
		c <- struct{}{}
	}(c, ci)

	println("NumGoroutine=", runtime.NumGoroutine())

	//读通道c,通过通道进行同步等待
	<-c

	//此时ci通道已经关闭，匿名函数启动的goroutine已经退出
	println("NumGoroutine=", runtime.NumGoroutine())

	//但此时通道还可以继续读取
	for v := range ci {
		println(v)
	}
}
