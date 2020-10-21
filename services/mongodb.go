package services

import (
	"fmt"
	"bitbucket.org/indoquran-api/config"
	"time"

	"gopkg.in/mgo.v2"
)

// MGOConnection defines the connection structure
type MGOConnection struct {
	session *mgo.Session
}

// MongoNewConnection handles connecting to a mongo database
func MongoNewConnection(dbName string) (conn *MGOConnection) {
	mongoHostPort := config.LoadConfig().Database.Host + ":" + config.LoadConfig().Database.Port
	info := &mgo.DialInfo{
		Addrs:    []string{mongoHostPort},
		Timeout:  60 * time.Second,
		Database: dbName,
		Username: config.LoadConfig().Database.Username,
		Password: config.LoadConfig().Database.Password,
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		fmt.Println("Host : " + mongoHostPort)
		fmt.Println("dbname : " + config.LoadConfig().Database.DatabaseName)
		fmt.Println("username: " + config.LoadConfig().Database.Username)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &MGOConnection{session}
	return conn
}

// MongoUse handles connect to a certain collection
func (conn *MGOConnection) MongoUse(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// MongoClose handles closing a database connection
func (conn *MGOConnection) MongoClose() {
	// This closes the connection
	conn.session.Close()
	return
}
