package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dags-/downloads/dl"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"os"
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

	go handleStop()

	panic(server.ListenAndServe(fmt.Sprintf(":%v", *port)))
}

func handleStop() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "stop" {
			fmt.Println("Stopping...")
			os.Exit(0)
			break
		}
	}
}