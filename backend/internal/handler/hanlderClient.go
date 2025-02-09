package handler

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addClient(c *gin.Context) {
	var client models.ClientHandler

	if err := c.BindJSON(&client); err != nil {
		logger.Log("Error", "Error binding JSON to struct:", err,
			fmt.Sprintf("Request Body: %s", c.Request.Body))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	clientData, err := h.services.AddClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error creating client",
		})
		return
	}

	c.JSON(http.StatusOK, clientData)
}

func (h *Handler) signIn(c *gin.Context) {
	var client models.ClientHandler

	if err := c.BindJSON(&client); err != nil {
		logger.Log("Error", "Error binding JSON to struct:", err,
			fmt.Sprintf("Request Body: %s", c.Request.Body))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	clientData, token, err := h.services.GenerateToken(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error creating client",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Token":    token,
		"Email":    clientData.Email,
		"Login":    clientData.Login,
		"ClientID": clientData.ClientID,
	})
}
