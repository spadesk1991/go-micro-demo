package lib

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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

func Call2(s selector.Selector, ctx context.Context, size int) {
	cli := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	p := map[string]int{"size": size}
	req := cli.NewRequest("prodService", "/news", &p)
	var res map[string]interface{}
	err := cli.Call(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res["data"])
}
