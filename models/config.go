package models

import (
	"indoquran-golang/config"
	"indoquran-golang/services"
)

// DatabaseName database name
var DatabaseName = config.LoadConfig().Database.DatabaseName

// DBConnect create a connection
var DBConnect = services.MGONewConnection(DatabaseName)
