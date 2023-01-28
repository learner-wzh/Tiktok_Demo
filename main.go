package main

import "tiktok_Demo/router_tiktok"

func main() {

	r := router_tiktok.InitRouter()

	err := r.Run()

	if err != nil {
		return
	}
}
