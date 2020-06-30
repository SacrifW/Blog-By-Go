package db

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

// DBConnection defines the connection structure
type DBConnection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(localhost string, BlogDB string) (conn *DBConnection) {
	info := &mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{localhost},
		// Timeout when a failure to connect to db
		Timeout: 60 * time.Second,
		// Database name
		Database: BlogDB,
		// Database credentials if your db is protected
		Username: os.Getenv("name"),
		Password: os.Getenv("password"),
	}

	session, err := mgo.DialWithInfo(info)


	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}


// Use handles connect to a certain collection
func (conn *DBConnection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *DBConnection) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
