package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
}
