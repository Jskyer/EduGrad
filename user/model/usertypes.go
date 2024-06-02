package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	Username string    `bson:"username,omitempty" json:"username,omitempty"`
	Password string    `bson:"password,omitempty" json:"password,omitempty"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Grade    string `bson:"grade,omitempty" json:"grade,omitempty"`
}
