package server

import (
	"github.com/Cracker-TG/line-notify-covid19-report/controllers"
	"github.com/gin-gonic/gin"
	//"github.com/vsouza/go-gin-boilerplate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	main := new(controllers.MainController)

	//router.GET("/ping", main.Status)

	v1 := router.Group("api/v1")
	{
		mainGroup := v1.Group("main")
		 {
			mainGroup.GET("/push-noti", main.PushNoti)
		}
	}
	
	return router
}