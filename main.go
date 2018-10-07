package main

import (
	server "github.com/RadicalPixels/server/server"
)

func main() {
	svr := server.NewServer(&server.Config{
		NodeURL:         "https://kovan.infura.io",
		ContractAddress: "0x2d31eb328000e3314243d49a459ae03127663ad0",
	})
	svr.Start()
}
