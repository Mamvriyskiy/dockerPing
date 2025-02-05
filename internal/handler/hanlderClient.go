package handler 

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
	"github.com/Mamvriyskiy/dockerPing/internal/services"
)

func (h *Handler) addPing(c *gin.Context) {
	var client models.ClientHandler

	if err := c.BindJSON(&client); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}

	_ = services.AddClient(client)
	
	// fmt.Println("+")
	// c.JSON(http.OK, map[string]interface{}{})
}

