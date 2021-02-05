package main

import (
	"family-board-api/auth"
	"family-board-api/config"
	"family-board-api/handler"
	"family-board-api/registry"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// envファイルを読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db初期化
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// DI
	repository := registry.NewRepository(db)
	userHandler := handler.NewUserHandler(repository)
	familyHandler := handler.NewFamilyHandler(repository)

	e := echo.New()

	// 認証
	e.POST("/login", userHandler.LoginWithLiff)

	// jwt認証が必要
	jwtConfig := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(jwtConfig))

	// family
	r.POST("/families", familyHandler.Create)
	r.POST("/families/join", familyHandler.Join)

	e.Logger.Fatal(e.Start(":8000"))
}
