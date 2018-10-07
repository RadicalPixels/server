package main

import (
	server "github.com/RadicalPixels/server/server"
)

func main() {
	svr := server.NewServer(&server.Config{
		NodeURL:         "https://kovan.infura.io",
		ContractAddress: "0xcfbded0bbf3726a056b1d9458308dd338e9eea63",
	})
	svr.Start()
}
