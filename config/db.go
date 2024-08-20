package config

import (
	"libro-electronico/helper/chicken"
	"libro-electronico/model"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGO_URI")

var mongoinfo = model.DBIngfo{
	DBString: MongoString,
	DBName:   "libroelectronico",
}

var Mongoconn, ErrorMongoconn = chicken.MongoConnect(mongoinfo)

var DB *mongo.Database