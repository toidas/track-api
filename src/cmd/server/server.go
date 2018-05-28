package main

import (
	"fmt"
	"../../routes/approuter"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	fmt.Println("Server run at port 8080")
	approuter.AppRouter(router)
	router.Run(":8080")
}
