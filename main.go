package main

import (
	"go-crowdfunding/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/go_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connected to mysql")

	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name, user.Email)
	// }

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

	// handler
	// service
	// repository
	// db
}
