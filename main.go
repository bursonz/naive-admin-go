package main

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/router"
	"time"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc
	app := gin.Default()
	config.Init()
	db.Init()
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	router.Init(app)

	if err := app.Run(":8800"); err != nil {
		panic(err)
	}
}
