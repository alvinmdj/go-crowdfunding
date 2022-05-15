package main

import (
	"fmt"
	"go-crowdfunding/auth"
	"go-crowdfunding/campaign"
	"go-crowdfunding/handler"
	"go-crowdfunding/helper"
	"go-crowdfunding/transaction"
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
	// setup database
	helper.LoadEnv()
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	// connect to database MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// setup auth
	authService := auth.NewService()

	// setup user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	// setup campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// setup transaction
	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// setup router
	router := gin.Default()

	// setup static file routes
	router.Static("/avatars", "./public/images/avatars")
	router.Static("/campaign-images", "./public/images/campaign-images")

	// setup api routes
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
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadCampaignImage)

	// transaction routes
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)

	router.Run()
}

// authMiddleware is a middleware function that checks if the user is authenticated.
// if user is authenticated, set current user data to context.
// if user is not authenticated, return status unauthorized & abort request.
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// check if authorization header is empty, if so, abort request
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			// as a middleware, abort the request to prevent further processing
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get token from authorization header
		tokenString := strings.Split(authHeader, " ")[1]

		// decode token, if error, abort request
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get claims from token, if failed, abort request
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get user id from claims & get user data from database
		// if failed, abort request
		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse(
				"Unauthorized", http.StatusUnauthorized, "error", nil,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set the current user to the context
		c.Set("currentUser", user)
	}
}
