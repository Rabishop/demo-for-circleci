package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"example.com/m/model"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取当前时间
	currentTime := time.Now()

	// 生成一个介于1到100之间的随机数
	randomNumber := rand.Intn(100) + 1

	// 构造响应消息
	message := fmt.Sprintf("Current time: %s\nRandom number: %d\nHello, World!", currentTime, randomNumber)

	// 将响应消息写入响应写入器
	fmt.Fprint(w, message)

	DB := model.MysqlConn()

	tx := DB.Begin()

	// get games by gameName
	if err := tx.Debug().Find("").Error; err != nil {
		log.Println(err)
		tx.Rollback()
	}

	tx.Commit()

}
