package handler

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addContainersStatus(c *gin.Context) {
	var containers []models.HistoryHandler
	if err := c.BindJSON(&containers); err != nil {
		logger.Log("Error", "Error binding JSON to struct:", err,
			fmt.Sprintf("Request Body: %s", c.Request.Body))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	err := h.services.AddContainersStatus(containers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error adding status",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) getContainersStatus(c *gin.Context) {
	id, ok := c.Get("clientID")
	if !ok {
		logger.Log("Warning", "Get clientID from context failed", nil)
		return
	}

	clientID, ok := id.(string)
	if !ok {
		logger.Log("Error", "Error converting value to string", nil, fmt.Sprintf("id = %s", id))

		return
	}

	containersData, err := h.services.GetContainersStatus(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error retrieving list of container status",
		})
		return
	}

	fmt.Println(containersData)
	
	c.JSON(http.StatusOK, containersData)
}

