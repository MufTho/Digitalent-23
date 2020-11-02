package main

import (
	"github.com/MufTho/Digitalent-23/app/controller"
	"github.com/gin-gonic/gin"
)

//Antrian @dec queue
func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("view/*") //load semua html yang ada didalam file view
	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrain/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}
