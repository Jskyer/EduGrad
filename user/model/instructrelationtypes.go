package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InstructRelation struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`

	StuID       string `bson:"stuId,omitempty" json:"stuId,omitempty"`
	StuName     string `bson:"stuName,omitempty" json:"stuName,omitempty"`
	TeacherID   string `bson:"teacherId,omitempty" json:"teacherId,omitempty"`
	TeacherName string `bson:"teacherName,omitempty" json:"teacherName,omitempty"`
	Permit      string `bson:"permit,omitempty" json:"permit,omitempty"`
	Title       string `bson:"title,omitempty" json:"title,omitempty"`
	State       string `bson:"state,omitempty" json:"state,omitempty"`
}
