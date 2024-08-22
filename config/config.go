package config

import (
	"libro-electronico/helper"
	"libro-electronico/helper/atdb"
	"libro-electronico/model"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var IPPort, Net = helper.GetAddress()

var PhoneNumber string

func SetEnv() {
	if ErrorMongoconn != nil {
		log.Println(ErrorMongoconn.Error())
	}
	profile, err := atdb.GetOneDoc[model.Profile](Mongoconn, "profile", primitive.M{})
	if err != nil {
		log.Println(err)
	}
	PublicKeyWhatsAuth = profile.PublicKey
	WAAPIToken = profile.Token
}