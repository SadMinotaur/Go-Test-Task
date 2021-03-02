package dao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
	typesF "medods/src/types"
)


const (
	mongoUrl   = "mongodb://localhost:27017/?readPreference=primary&ssl=false"
	database   = "medods"
	rTokensCol = "tokens"
)

func SaveToken(rt typesF.RToken) error {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		return err
	}
	collection := session.DB(database).C(rTokensCol)
	defer session.Close()
	_, err = collection.RemoveAll(bson.M{"GUID": rt.GUID})
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