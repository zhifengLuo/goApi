package main

import "goapi/router"

func main() {
	r := router.NewRouter()
	r.Run(":8088")
}
