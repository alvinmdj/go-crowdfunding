package main

import (
	"fmt"
	"go-crowdfunding/auth"
	"go-crowdfunding/campaign"
	"go-crowdfunding/handler"
	"go-crowdfunding/helper"
	"go-crowdfunding/payment"
	"go-crowdfunding/transaction"
	"go-crowdfunding/user"
	webHandler "go-crowdfunding/web/handler"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
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

	// setup payment
	paymentService := payment.NewService()

	// setup transaction
	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// setup user web
	userWebHandler := webHandler.NewUserHandler()

	// setup router
	router := gin.Default()

	// setup cors middleware
	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
  }))

	// setup templates from /web/templates folder
	router.HTMLRender = loadTemplates("./web/templates")

	// setup static file routes
	router.Static("/avatars", "./public/images/avatars")
	router.Static("/campaign-images", "./public/images/campaign-images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	// user web routes
	router.GET("/users", userWebHandler.Index)

	// setup api routes
	api := router.Group("/api/v1")

	// user routes
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	// campaign routes
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignDetails)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadCampaignImage)

	// transaction routes
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run() // default port 8080
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

// loadTemplates is a function to load HTML templates
func loadTemplates(templatesDir string) multitemplate.Renderer {
  r := multitemplate.NewRenderer()

	// load all file from /layouts directory
  layouts, err := filepath.Glob(templatesDir + "/layouts/*")
  if err != nil {
    panic(err.Error())
  }

	// load all folders from /templates directory
  includes, err := filepath.Glob(templatesDir + "/**/*")
  if err != nil {
    panic(err.Error())
  }

  // Generate our templates map from our layouts/ and includes/ directories
  for _, include := range includes {
    layoutCopy := make([]string, len(layouts))
    copy(layoutCopy, layouts)
    files := append(layoutCopy, include)
    r.AddFromFiles(filepath.Base(include), files...)
  }
  return r
}
