package handler 

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
	//"github.com/Mamvriyskiy/dockerPing/internal/services"
)

func (h *Handler) addClient(c *gin.Context) {
	var client models.ClientHandler

	if err := c.BindJSON(&client); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	clientData, err := h.services.AddClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка создания клиента",
		})
		return
	}
	
	// fmt.Println("+")
	c.JSON(http.StatusOK, clientData)
}
