package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("Ping")
}

func (*St) Pang() {
	println("Pang")
}

//func (St) String() {
//	println("111")
//}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	//i没有实现接口Anter，所有程序不执行这里
	if p, ok := i.(Anter); ok {
		p.String()
	}

	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name)
	}
}
