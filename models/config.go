package models

import (
	"indoquran-golang/config"
	"indoquran-golang/db"
)

// Database name
var databaseName = config.LoadConfig().Database.DatabaseName

// Create a connection
var dbConnect = db.NewConnection(databaseName)
