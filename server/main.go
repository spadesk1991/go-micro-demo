package main

import (
	"net/http"
	"strconv"

	"github.com/spadesk1991/go-micro-demo/model"

	"github.com/gin-gonic/gin"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs("192.168.1.101:8500"))
	//reg:= consul
	ginRouter := gin.Default()
	ginRouter.GET("/user", func(context *gin.Context) {
		context.String(200, "ok")
	})
	type ProdModel struct {
		ID   int
		Name string
	}
	ginRouter.POST("/news", func(context *gin.Context) {
		var p model.ProdParams
		context.ShouldBind(&p)
		count := p.Size
		data := make([]ProdModel, count)
		for i := 0; i < count; i++ {
			data[i] = ProdModel{1000 + i, "name_" + strconv.Itoa(1000+i)}
		}
		context.JSON(http.StatusOK, gin.H{"data": data})
	})
	server := web.NewService(
		//web.Address(":8001"),
		web.Name("prodService"),
		web.Handler(ginRouter),
		web.Registry(reg),
	)
	server.Init()
	server.Run()
	//ginRouter.Run(":8000")
}
