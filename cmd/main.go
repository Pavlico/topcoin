package main

import "github.com/Pavlico/topcoin/internal/http/server"

func main() {
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
