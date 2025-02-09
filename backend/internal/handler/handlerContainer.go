package handler

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addContainer(c *gin.Context) {
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

	var container models.ContainerHandler
	if err := c.BindJSON(&container); err != nil {
		logger.Log("Error", "Error binding JSON to struct:", err, fmt.Sprintf("Request Body: %s", c.Request.Body))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	containerData, err := h.services.AddContainer(container, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error adding container IP",
		})
		return
	}

	c.JSON(http.StatusOK, containerData)
}

func (h *Handler) getContainers(c *gin.Context) {
	containersData, err := h.services.GetContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Error retrieving list of container IPs",
		})
		return
	}

	c.JSON(http.StatusOK, containersData)
}
