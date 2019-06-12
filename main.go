package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	router.GET("/user", inDB.GetUsers)
	router.GET("/user/:id", inDB.GetUser)
	router.POST("/user", inDB.CreateUser)
	router.PUT("/user/:id", inDB.UpdateUser)
	router.DELETE("/user/:id", inDB.DeleteUser)
	router.Run(":8089")
}
