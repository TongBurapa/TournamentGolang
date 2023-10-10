package Tourcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTournamentController(c *gin.Context) error {

	return c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("test"))
}
