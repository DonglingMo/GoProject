package main

import (
	"fmt"
	"math"
)

type MyStruct struct {
	X, Y float64
}

type MyFloat float64

// 类型方法
// 非指针类型ms为MyStruct副本
// 指针类型原始数据修改，避免大结构体复制
func (ms MyStruct) Distance() float64 {
	return ms.X*ms.X + ms.Y*ms.Y
}

func Abs(v MyStruct) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 数值类型方法，类型定义和方法需要在同一个包
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyStruct2 struct {
	X, Y float64
}

func (v MyStruct2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *MyStruct2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsF(ms2 MyStruct2) float64 {
	return math.Sqrt(ms2.X*ms2.X + ms2.Y*ms2.Y)
}

func Scale(ms2 *MyStruct2, f float64) {
	ms2.X = ms2.X * f
	ms2.Y = ms2.Y * f
}

func main() {
	p := MyStruct{3, 4}
	fmt.Println(p.Distance()) // 调用类型方法、
	fmt.Println(p)
	// 方法结构体作为函数参数
	fmt.Println(Abs(p))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v := MyStruct2{3, 4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())

	v2 := MyStruct2{3, 4}
	// 方法function调用，接受指针类型时必须传入指针
	// 类型method定义为指针是，可以传入一个普通的非指针变量或者可以接受一噩耗指针
	// v.Scale(5) 解释为 (&v).Scale(5)

	// 如果是普通指定的类型，普通方法则传入对应类型
	// 类类型方法也是既可以接受普通类型，又可以接受指针
	// 方法调用 p.Abs() 会被解释为 (*p).Abs()
	// 建议是所有的类型方法 都应该定义为指针
	Scale(&v2, 10)
	fmt.Println(AbsF(v2))

}
