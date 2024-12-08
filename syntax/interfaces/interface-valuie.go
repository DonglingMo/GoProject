package main

import (
	"fmt"
	"math"
)

type II interface {
	MM()
}

type TT struct {
	S string
}

func (tt *TT) MM() {
	if tt == nil {
		fmt.Println("TT is nil")
		return
	}
	fmt.Println(tt.S)
}

type F float64

func (f F) MM() {
	fmt.Println(f)
}

func main() {
	var ii II
	// 接口没有实现时调用会painc,没有实现的接口不会包含任何信息
	// describe(ii)
	// ii.MM()

	ii = &TT{"Hello"}
	describe(ii)
	ii.MM()

	ii = F(math.Pi)
	describe(ii)
	ii.MM()

	var tt *TT
	ii = tt
	describe(ii)
	// 接口为空时仍然可以调用， nil的具体值保存，接口本身不为nil
	ii.MM()
}

func describe(i II) {
	fmt.Printf("(%v, %T)\n", i, i)
}
