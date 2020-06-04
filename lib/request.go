package lib

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	models_protos "github.com/spadesk1991/go-micro-demo/server/models"

	"github.com/micro/go-micro/client"

	"github.com/micro/go-micro/client/selector"
	myhttp "github.com/micro/go-plugins/client/http"
)

func Call(method, url, params string, body io.Reader) (buff []byte, err error) {
	request, err := http.NewRequest(method, url+params, body)
	if err != nil {
		return
	}
	client := http.DefaultClient
	r, err := client.Do(request)
	defer r.Body.Close()
	if err != nil {
		return
	}
	buff, _ = ioutil.ReadAll(r.Body)
	return
}

func Call2(s selector.Selector, ctx context.Context, size int) (res models_protos.ProdListResponse) {
	cli := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	p := models_protos.ProdsRequest{Size: int32(size)}
	req := cli.NewRequest("prodService", "/news", &p)
	err := cli.Call(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetData())
	return
}
