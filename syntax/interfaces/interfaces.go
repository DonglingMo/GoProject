package main

import (
	"fmt"
	"math"
)

type Abser interface {
	AbsImpl() float64
}

func main() {
	var abser Abser
	f := MyFf(-math.Sqrt2)
	v := MyPointer{3, 4}

	// MyFF类型实现接口
	abser = f
	// *MyPointer类型实现接口
	abser = &v
	// MyPointer没有实现接口
	///abser = v

	fmt.Println(abser.AbsImpl())
}

type MyFf float64

func (f MyFf) AbsImpl() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyPointer struct {
	X, Y float64
}

func (mp *MyPointer) AbsImpl() float64 {
	return math.Sqrt(mp.X*mp.X + mp.Y*mp.Y)
}
