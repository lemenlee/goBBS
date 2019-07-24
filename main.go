package main

import "bbs/api"

func main() {
	r := api.SetupRouter()
	r.Run()
}
