package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取当前时间
	currentTime := time.Now()

	// 输出当前时间
	fmt.Println("Current time:", currentTime)

	// 生成一个介于1到100之间的随机数
	randomNumber := rand.Intn(100) + 1

	// 输出随机数
	fmt.Println("Random number:", randomNumber)

	// 输出 "Hello, World!"
	fmt.Println("Hello, World!")
}
