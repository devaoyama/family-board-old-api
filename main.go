package main

import (
	"family-board-api/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// envファイルを読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = config.InitDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("完了")
}
