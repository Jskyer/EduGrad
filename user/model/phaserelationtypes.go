package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhaseRelation struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	UserId   string `bson:"userId,omitempty" json:"userId,omitempty"`
	PhaseId  string `bson:"phaseId,omitempty" json:"phaseId,omitempty"`
	Identity string `bson:"identity,omitempty" json:"identity,omitempty"`
	Pass     int64  `bson:"pass,omitempty" json:"pass,omitempty"`
}
