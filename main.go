package main

import (
	"flag"
	"fmt"
	"github.com/dags-/downloads/dl"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	port := flag.Int("port", 8083, "The server port")

	config := dl.LoadConfig()
	cache := dl.NewCache(config)

	router := routing.New()
	router.Get("/<repo>/<id>", func(c *routing.Context) (error) {
		repo := c.Param("repo")
		id := c.Param("id")
		url, err := cache.Get(repo, id)
		if err == nil {
			c.Redirect(url, 301)
		}
		return err
	})

	server := fasthttp.Server{
		Handler:            router.HandleRequest,
		GetOnly:            true,
		DisableKeepalive:   true,
		ReadTimeout:        time.Duration(time.Second * 2),
		WriteTimeout:       time.Duration(time.Second * 2),
		MaxConnsPerIP:      3,
		MaxRequestsPerConn: 1,
		MaxRequestBodySize: 0,
	}

	panic(server.ListenAndServe(fmt.Sprintf(":%v", *port)))
}
