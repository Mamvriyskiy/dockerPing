package handler 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

func (h *Handler) addContainer(c * gin.Context) {
	// id, ok := c.Get("userId")
		// if !ok {
	// 	logger.Log("Warning", "Get", "Error get userID from context", nil, id)
	// 	return
	// }

	var container models.ContainerHandler
	if err := c.BindJSON(&container); err != nil {
		logger.Log("Error", "c.BindJSON()", "Error bind json:", err, "")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	containerData, err := h.services.AddContainer(container)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errors": "Ошибка при добавлении IP контейнера",
		})
		return
	}

	c.JSON(http.StatusOK, containerData)
}

// func (h *Handler) deleteContainer(c *gin.Context) {

// }
