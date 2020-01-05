package main

import (
				  "github.com/gin-gonic/gin"
)

func APIRouter() *gin.Engine{

				  router:=gin.Default()
				  router.GET("compuer", func(c *gin.Context) {
					c.JSON(200,gin.H{
									  "message":"北京车音智能科技有限公司",
					})
				  })
				
				  return router

}
