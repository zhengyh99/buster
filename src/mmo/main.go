package main

import "buster/bnet"

func main() {
	server := bnet.NewServer()
	
	server.Run()
}
