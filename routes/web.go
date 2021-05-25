package routes

import (
	. "github.com/ArtisanCloud/go-framework/app/http/controllers"
	. "github.com/ArtisanCloud/go-framework/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeWebRoutes(router *gin.Engine) {

	//println("%T", router)

	apiRouter := router.Group("/")
	{
		apiRouter.Use(Maintenance, AuthAPI, AuthWeb)
		{
			// Handle the index route
			apiRouter.GET("/", WebGetHome)
			//apiRouter.POST("/make", ValidateRequestMakeWelcome, ctlWelcome.APIMakeWelcome)
			// apiRouter.PUT("/somePut", putting)
			// apiRouter.DELETE("/someDelete", deleting)
			// apiRouter.PATCH("/somePatch", patching)
			// apiRouter.HEAD("/someHead", head)
			// apiRouter.OPTIONS("/someOptions", options)

		}
	}

}
