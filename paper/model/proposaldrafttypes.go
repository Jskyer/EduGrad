package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProposalDraft struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	InstructId string `bson:"instructId,omitempty" json:"instructId,omitempty"`
	Content    string `bson:"content,omitempty" json:"content,omitempty"`
	Comment    string `bson:"comment,omitempty" json:"comment,omitempty"`
}
