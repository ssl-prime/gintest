package main

import (
	"gintest/api/controller"
	"gintest/api/util"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/aicumen/v1")
	{
		v1.POST("/insertinfo", controller.InsertInfo)
		v1.PUT("/updateInfo", controller.UpdateInfo)
		v1.POST("/deleteInfo", controller.DeleteInfo)
		v1.GET("/getkey", controller.GetKey)
	}
	router.Run(":8090")
}
func init() {
	util.ConnectOrbitDB("aicumen")
}
