package config

import (
	"log"

	"libro-electronico/helper"
	"libro-electronico/helper/chicken"
	"libro-electronico/model"

	"gopkg.in/mgo.v2/bson"
)

var IPPort, Net = helper.GetAddress()

var PhoneNumber string


func SetEnv() {
	if ErrorMongoconn != nil {
		log.Println(ErrorMongoconn.Error())
	}

	// Mengambil dokumen pertama dari koleksi "profile"
	profile, err := chicken.GetOneDoc[model.Profile](Mongoconn, "profile", bson.M{})
	if err != nil {
		log.Println(err)
		return
	}

	// Set nilai-nilai dari profile yang diambil
	PublicKeyWhatsAuth = profile.PublicKey
	WAAPIToken = profile.Token
}
