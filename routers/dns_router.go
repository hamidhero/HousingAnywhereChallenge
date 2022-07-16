package routers

import (

	"github.com/gin-gonic/gin"
	"housingAnywherChallenge/controllers"
)

func GetRouter() (router *gin.Engine) {
	router = gin.Default()

	//generate prefix /api
	api := router.Group("api")
	{
		api.POST("location", controllers.GetDataBankLocation)
	}

	return
}
