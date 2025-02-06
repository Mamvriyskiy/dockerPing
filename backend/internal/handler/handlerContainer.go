package handler 

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
)

func (h *Handler) addContainer(c * gin.Context) {
	id, ok := c.Get("clientID")
	if !ok {
		logger.Log("Warning", "Get", "Error get clientID from context", nil, id)
		return
	}

	clientID, ok := id.(string)
	if !ok {
		logger.Log("Error", "Conversion", "Error conversion any to string", nil, id)
		return
	}

	var container models.ContainerHandler
	if err := c.BindJSON(&container); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	containerData, err := h.services.AddContainer(container, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка при добавлении IP контейнера",
		})
		return
	}

	c.JSON(http.StatusOK, containerData)
}

func (h *Handler) getContainers(c *gin.Context) {
	containersData, err := h.services.GetContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка при получении списка IP контейнеров",
		})
		return
	}

	fmt.Println(containersData)
	c.JSON(http.StatusOK, containersData)
}
