package services

import (
	"fmt"
	"indoquran-golang/config"
	"time"

	"gopkg.in/mgo.v2"
)

// MGOConnection defines the connection structure
type MGOConnection struct {
	session *mgo.Session
}

// MGONewConnection handles connecting to a mongo database
func MGONewConnection(dbName string) (conn *MGOConnection) {
	info := &mgo.DialInfo{
		Addrs:    []string{config.LoadConfig().Database.Host},
		Timeout:  60 * time.Second,
		Database: dbName,
		Username: config.LoadConfig().Database.Username,
		Password: config.LoadConfig().Database.Password,
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		fmt.Println("Host : " + config.LoadConfig().Database.Host)
		fmt.Println("dbname : " + config.LoadConfig().Database.DatabaseName)
		fmt.Println("username: " + config.LoadConfig().Database.Username)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &MGOConnection{session}
	return conn
}

// MGOUse handles connect to a certain collection
func (conn *MGOConnection) MGOUse(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// MGOClose handles closing a database connection
func (conn *MGOConnection) MGOClose() {
	// This closes the connection
	conn.session.Close()
	return
}
