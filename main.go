package main

import (
	"indoquran-golang/config"
	"indoquran-golang/services"
)

// init gets called before the main function
func init() {
	config.SetLogger()
}

func main() {
	services.StartServer()
}
