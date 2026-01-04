package routes

import (
	"Zeus/api"

	"github.com/gin-gonic/gin"
)

func V1(engine *gin.Engine) {
	v1 := engine.Group("api/v1")
	{
		api.User.API(v1)
	}
}
