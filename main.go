package main

import (
	"canvine/router"
)

func main() {
	r := router.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}