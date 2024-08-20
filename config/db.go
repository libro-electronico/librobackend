package config

import (
	"libro-electronico/helper/chicken"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRINGOLD")

var mongoinfo = chicken.DBIngfo{
	DBString: MongoString,
	DBName:   "libroelectronico",
}

var Mongoconn, ErrorMongoconn = chicken.MongoConnect(mongoinfo)

var DB *mongo.Database