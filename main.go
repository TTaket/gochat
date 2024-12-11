package main

import "github.com/TTaket/gochat/router"

func main() {
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
