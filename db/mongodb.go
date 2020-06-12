package db

import (
	"fmt"
	"indoquran-golang/config"
	"time"

	"gopkg.in/mgo.v2"
)

// Connection defines the connection structure
type Connection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(dbName string) (conn *Connection) {
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
	conn = &Connection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *Connection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *Connection) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
