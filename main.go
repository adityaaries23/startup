package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	RepoUser := user.NewRepository(db)
	ServiceUser := user.NewService(RepoUser)
	HandlerUser := handler.NewUserHandler(ServiceUser)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", HandlerUser.RegisterUser)
	router.Run()
}
