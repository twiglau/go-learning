package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"apiserver/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")

	if err := http.ListenAndServe(":8080", g); err != nil {
		log.Fatal(err)
	}
	// err := http.ListenAndServe(":8080", g)
	// 方式一：
	// log.Println(err)
	// 方式二：
	// log.Printf("%v", err)
	// 方式三：
	// log.Fatal(http.ListenAndServe(":8080", g),)
	// 不推荐
	// log.Printf(http.ListenAndServe(":8080", g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
