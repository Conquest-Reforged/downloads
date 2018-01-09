package main

import (
	"time"
	"github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
	"fmt"
)

func main() {
	router := routing.New()
	router.Get("/<repo>/<id>", func(c *routing.Context) error {

		return nil
	})

	server := fasthttp.Server{
		Handler:            router.HandleRequest,
		GetOnly:            false,
		DisableKeepalive:   true,
		ReadBufferSize:     10240,
		WriteBufferSize:    25600,
		ReadTimeout:        time.Duration(time.Second * 2),
		WriteTimeout:       time.Duration(time.Second * 2),
		MaxConnsPerIP:      3,
		MaxRequestsPerConn: 1,
		MaxRequestBodySize: 0,
	}

	panic(server.ListenAndServe(fmt.Sprintf(":%v", 8083)))
}
