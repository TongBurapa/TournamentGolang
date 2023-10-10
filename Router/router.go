package router

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/getTournament", tourAPI.getTournamentController)

}
