package config

import (
	"libro-electronico/helper"
	"libro-electronico/model"
	"os"
)

var MongoString string = os.Getenv("MONGO_URI")

var mongoinfo = model.DBIngfo{
	DBString: MongoString,
	DBName:   "libroelectronico",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
