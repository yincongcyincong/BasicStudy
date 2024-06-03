package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/yincongcyincong/BasicStudy/httpserver/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// login
	router.LoadHTMLGlob("view/**/*")
	router.StaticFS("/static", gin.Dir("./static", true))

	router.PUT("/user", controller.AddStudyUser)
	router.POST("/user", controller.UpdateStudyUser)
	router.DELETE("/user", controller.DelStudyUser)
	router.PUT("/question", controller.AddStudyQuestion)
	router.POST("/choose_user", controller.ChooseUser)

	return router
}
