package id

import (
	"gopkg.in/mgo.v2/bson"
)

type ID bson.ObjectId

func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

func NewID() ID {
	return StringToID(bson.NewObjectId().Hex())
}
