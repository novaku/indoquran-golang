package main

import (
	"indoquran-golang/config"
	"indoquran-golang/server"
)

// init gets called before the main function
func init() {
	config.SetLogger()
}

// main function
func main() {
	server.StartServer()
}
