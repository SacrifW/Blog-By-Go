package models

import (
	"os"
	"Go-Blog-Project-master/db"

)

var server = os.Getenv("localhost")

var databaseName = os.Getenv("BlogDB")


var dbConnect = db.NewConnection(server, databaseName)
