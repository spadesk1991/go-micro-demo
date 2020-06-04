package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/spadesk1991/go-micro-demo/lib"

	"github.com/micro/go-micro/client/selector"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	getServicd, err := reg.GetService("prodService")
	if err != nil {
		log.Fatal(err)
	}
	//myhttp.NewClient(client.)
	var s selector.Selector
	// 定义选择器类型
	s = selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)
	router := gin.Default()

	type resModel struct {
		ID   int
		Name string
	}
	type res struct {
		Data []resModel `json:"data"`
	}

	router.GET("/news", func(ctx *gin.Context) {
		next := selector.RoundRobin(getServicd)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		r := new(res)
		buff, err := lib.Call("POST", "http://"+node.Address+"/news", "", nil)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(buff, r)
		ctx.JSON(http.StatusOK, gin.H{"datas": r})
	})

	router.GET("/v2/news", func(ctx *gin.Context) {
		count, _ := ctx.GetQuery("count")
		size, _ := strconv.Atoi(count)
		res := lib.Call2(s, ctx, size)
		ctx.JSON(http.StatusOK, gin.H{"data": res.GetData()})
	})

	srv := web.NewService(
		web.Name("web_service"),
		web.Address(":8003"),
		web.Registry(reg),
		web.Handler(router),
	)
	srv.Run()
}
