package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
}
