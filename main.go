package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter() *gin.Engine {
	api := gin.Default()
	api.Use(gin.Recovery())

	Router := api.Group("/api/v1")
	{
		Router.GET("/consume-date")
	}

	return api
}

func main() {
	r := initRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("Listen err %v :", err)
		}
	}()

}
