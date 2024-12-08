package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go 运行的系统环境：")
	// 类似的判断条件之前可以执行一段表达式
	// 不需要break
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}

	t := time.Now()
	// 无条件switch
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
