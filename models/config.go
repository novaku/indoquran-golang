package models

import (
	"indoquran-golang/config"
	"indoquran-golang/services"
)

// Database name
var databaseName = config.LoadConfig().Database.DatabaseName

// Create a connection
var dbConnect = services.MGONewConnection(databaseName)
