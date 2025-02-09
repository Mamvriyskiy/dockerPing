package handler

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/services"
	"github.com/Mamvriyskiy/dockerPing/logger"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

const (
	signingKey  = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
	pingerToken = "hsHcmJkmHaJIUzUxMiIsInR5cC3jhmdHJ7H.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJKb2huIEdvbGQiLCJhZG1pbiI6dHJ1ZX0K.LIHjWCBORSWMEibq-tnT8ue_deUqZx1K0XxCOXZRrBI"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		var err error
		if tokenString == "" {
			tokenString, err = c.Cookie("jwt")
			if err != nil {
				logger.Log("Error", "Error retrieving JWT from cookies", err)
			}
		}

		if tokenString == "" {
			logger.Log("Error", "Empty token in Authorization header and cookies", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty token"})
			c.Abort()
			return
		}

		if tokenString == pingerToken {
			logger.Log("Info", "Using pinger token for authentication", nil)
			c.Next()
		} else {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(signingKey), nil
			})

			if err != nil {
				logger.Log("Error", "Error parsing JWT", err, fmt.Sprintf("token = %s", tokenString))
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "detail:": err.Error()})
				c.Abort()
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				logger.Log("Info", "Token valid, setting clientID", nil, fmt.Sprintf("clientID = %s", claims["clientID"]))
				c.Set("clientID", claims["clientID"])
				c.Next()
			} else {
				logger.Log("Error", "Invalid JWT token", err, fmt.Sprintf("token = %s", tokenString))
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
		}
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(ctx *gin.Context) {
		fmt.Println("Requested URL:", ctx.Request.URL.String())
		fmt.Println("Request Method:", ctx.Request.Method)
		ctx.Next()
	})

	router.Use(AuthMiddleware())

	auth := router.Group("/auth")
	auth.POST("/sign-up", h.addClient)
	auth.POST("/sign-in", h.signIn)

	api := router.Group("/api")
	api.POST("/ping", h.addContainer)
	api.GET("/ping", h.getContainers)

	api.POST("/pinger", h.addContainersStatus)

	return router
}
