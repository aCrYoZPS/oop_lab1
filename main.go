package main

import "oopLab1/core"

func main() {
	server := core.NewEchoServer()
	server.Start()
}
