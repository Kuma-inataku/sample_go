package main

import (
	"gin-api-sample/internal/infrastructure/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
