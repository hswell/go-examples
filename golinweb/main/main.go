package main

import (
	"golinweb/golin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() golin.HandlerFunc {
	return func(c *golin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := golin.New()
	r.Use(golin.Logger()) // global midlleware
	r.GET("/", func(c *golin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello golin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *golin.Context) {
			// expect /hello/golinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}

