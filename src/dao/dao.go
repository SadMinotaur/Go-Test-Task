package dao

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	typesF "medods/src/types"
)


const (
	mongoUrl   = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	database   = "medods"
	rTokensCol = "tokens"
)

func saveToken(rt typesF.RToken) error {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		return err
	}
	collection := session.DB(database).C(rTokensCol)
	session.Close()
	info, err := collection.RemoveAll(bson.M{"GUID" : rt.GUID})
	if err != nil {
		return err
	}
	bcryptToken, err := bcrypt.GenerateFromPassword([]byte(rt.Token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	rt.Token = string(bcryptToken)
	err = collection.Insert(rt)
	if err != nil {
		return err
	}
	return nil
}