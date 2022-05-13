/*
URL的規則與匹配路徑，以及該動作後續的執行行為
*/
package config

import (
	// "backend/app"

	"fmt"
	"thundermeet_backend/app/controller"

	"github.com/gin-gonic/gin"
)

func RouteUsers(r *gin.Engine) {
	posts := r.Group("/v1/users")
	{
		fmt.Print("in router")
		posts.POST("/", controller.NewUsersController().CreateUser)
		posts.POST("/login", controller.NewUsersController().Login)
		posts.GET("/", controller.QueryUsersController().CheckUser)
		posts.PATCH("/", controller.UpdateUsersController().UpdateUserInfo)
		posts.PATCH("/resetPassword", controller.UpdateUsersController().ResetPassword)
	}

	events := r.Group("/v1/events")
	{
		events.POST("/", controller.CreateEventsController().CreateEvent)
		events.GET("/:event_id", controller.GetEventsController().GetEvent)
		events.GET("/", controller.GetEventsController().GetEvents)
		events.PATCH("/", controller.UpdateEventsController().UpdateEvent)
	}

	timeblocks := r.Group("/v1/timeblocks")
	{
		timeblocks.POST("/", controller.CreateTimeblocksController().CreateTimeblock)
	}
}
