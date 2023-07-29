package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"GitHubWebHookDemo/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Use(gin.Logger())
	if err := r.Run(":80"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
