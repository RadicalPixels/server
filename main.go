package main

import (
	server "github.com/RadicalPixels/server/server"
)

func main() {
	svr := server.NewServer(&server.Config{
		NodeURL:         "https://kovan.infura.io",
		ContractAddress: "0xa74e7fea1db19f0f41d054854f4d950f1c6ff513",
	})
	svr.Start()
}
