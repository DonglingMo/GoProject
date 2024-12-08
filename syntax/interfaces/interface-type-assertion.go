package main

import "fmt"

func main() {
	var i interface{} = "hello world"
	// 接口类型断言,获取接口的底层数据类型赋值给s
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// 不存在赋值零值，false
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64)
	// fmt.Println(f)

	do(21)
	do("hello")
	do(true)
}

// 类型判断
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 长度为 %v 字节\n", v, len(v))
	default:
		fmt.Printf("我不知道类型 %T!\n", v)
	}
}
