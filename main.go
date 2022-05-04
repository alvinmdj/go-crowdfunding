package main

import (
	"go-crowdfunding/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Save from service"
	userInput.Occupation = "Software Engineer"
	userInput.Email = "test@google.com"
	userInput.Password = "12345678"

	userService.RegisterUser(userInput)

	// input from user
	// handler : mapping user input to input struct
	// service : mapping from input struct to User struct
	// repository
	// db
}
