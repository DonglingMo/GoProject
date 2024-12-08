package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 类型+接口方法隐式实现
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M() // 类型断言，i 转换为 T 类型，然后执行 M()
}
