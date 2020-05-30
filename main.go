package main

import (
	"indoquran-golang/config"
	"indoquran-golang/services"
)

var conf *config.Config

// init gets called before the main function
func init() {
	// conf = config.LoadConfig()
}

func main() {
	services.StartServer()
}
