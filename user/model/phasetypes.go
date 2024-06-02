package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phase struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	Term    string `bson:"term,omitempty" json:"term,omitempty"`
	Process int64  `bson:"process,omitempty" json:"process,omitempty"`
}
