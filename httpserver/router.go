package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/yincongcyincong/BasicStudy/httpserver/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	//authorized := router.Group("/")
	//authorized.POST("/login", loginEndpoint)
	//authorized.POST("/submit", submitEndpoint)
	//authorized.POST("/read", readEndpoint)

	router.PUT("/user", controller.AddStudyUser)
	router.POST("/user", controller.UpdateStudyUser)
	router.DELETE("/user", controller.DelStudyUser)

	router.PUT("/question", controller.AddStudyQuestion)

	return router
}
