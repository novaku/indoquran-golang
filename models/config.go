package models

import (
	"bitbucket.org/indoquran-api/config"
	"bitbucket.org/indoquran-api/services"
)

// DatabaseName database name
var DatabaseName = config.LoadConfig().Database.DatabaseName

// DBConnect create a connection
var DBConnect = services.MongoNewConnection(DatabaseName)
