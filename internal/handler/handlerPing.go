package handler 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addPing(c *gin.Context) {
	var models.ClientHandler
	
	if err := c.BindJSON(&client); err != nil {
		//log
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}

	fmt.Println("+")
}

