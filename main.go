package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs(":8500"))

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	})
	srv := web.NewService(
		web.Name("prodService"),
		web.Address(":8001"),
		web.Handler(router),
		web.Registry(reg),
	)
	srv.Run()
}
