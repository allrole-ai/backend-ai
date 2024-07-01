package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap     string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
	Salt            string             `bson:"salt,omitempty" json:"salt,omitempty"`
}

type Password struct {
	Password        string `bson:"password,omitempty" json:"password,omitempty"`
	Newpassword     string `bson:"newpass,omitempty" json:"newpass,omitempty"`
	Confirmpassword string `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
}
