package handler

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

const signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"

// Middleware для извлечения данных из JWT и добавления их в контекст запроса.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверить URL запроса
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			// Если URL неначинается с /api, пропустить проверку JWT
			c.Next()
			return
		}

		// Получить токен из заголовка запроса или из куки
		tokenString := c.GetHeader("Authorization")
		fmt.Println("Token:", tokenString)
		var err error
		if tokenString == "" {
			// Если токен не найден в заголовке, попробуйте из куки
			tokenString, err = c.Cookie("jwt")
			if err != nil {
				logger.Log("Error", "c.Cookie(jwt)", "Error", err, "jwt")
			}
		}

		// Проверить, что токен не пустой
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty token"})
			c.Abort()
			return
		}

		// Парсинг токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Здесь нужно вернуть ключ для проверки подписи токена.
			// В реальном приложении, возможно, это будет случайный секретный ключ.
			return []byte(signingKey), nil
		})
		
		// fmt.Println(err) 
		// Проверить наличие ошибок при парсинге токена
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "detail:": err.Error()})
			c.Abort()
			return
		}

		// Добавить данные из токена в контекст запроса
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("+")
			c.Set("userId", claims["userID"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(AuthMiddleware())

	auth := router.Group("/auth")
	auth.POST("/sign-up", h.addClient)
	auth.POST("/sign-in", h.signIn)

	api := router.Group("/ping")
	api.POST("/", h.addContainer)


	return router
}

