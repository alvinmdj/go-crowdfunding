package main

import (
	"fmt"
	"go-crowdfunding/auth"
	"go-crowdfunding/campaign"
	"go-crowdfunding/handler"
	"go-crowdfunding/helper"
	"go-crowdfunding/user"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	helper.LoadEnv()
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()

	router.Static("/avatars", "./public/images/avatars")

	api := router.Group("/api/v1")

	// user routes
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// campaign routes
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignDetails)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			// as a middleware, abort the request to prevent further processing
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user) // set the current user to the context
	}
}
