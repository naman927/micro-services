package main

import (
	"github.com/gin-gonic/gin"
	"github.com/store/controller"
)

func main() {

	r := gin.Default()

	r.POST("/upsert", controller.APIUpsertProduct)

	r.Run(":8081")
}
