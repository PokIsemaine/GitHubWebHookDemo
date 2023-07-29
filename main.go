package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"GitHubWebHookDemo/routers"
)

// https://www.fenghong.tech/blog/go/gin-use-in-webhook/

func main() {
	r := routers.SetupRouter()
	r.Use(gin.Logger())
	if err := r.Run(":80"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
