package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Certificate struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	Cname            string        `bson:"cname" json:"cname"`
	SigningAlgorithm string        `bson:"signingalgorithm" json:"signingalgorithm"`
	Issuer           string        `bson:"issuer" json:"issuer"`
	NotAfter         string        `bson:"notafter" json:"notafter"`
	NotBefore        string        `bson:"notbefore" json:"notbefore"`
	ServerId         bson.ObjectId `bson:"serverid"`
	ServerName       string        `bson:"serverName" json:"servername"`
}

// func (c *certificate) TimeToExpiry() {
// 	expiry := time.Now().Sub(c.NotAfter)
// }
