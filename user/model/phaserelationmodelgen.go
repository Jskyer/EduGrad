// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type phaseRelationModel interface {
	Insert(ctx context.Context, data *PhaseRelation) error
	FindOne(ctx context.Context, id string) (*PhaseRelation, error)
	Update(ctx context.Context, data *PhaseRelation) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultPhaseRelationModel struct {
	conn *mon.Model
}

func newDefaultPhaseRelationModel(conn *mon.Model) *defaultPhaseRelationModel {
	return &defaultPhaseRelationModel{conn: conn}
}

func (m *defaultPhaseRelationModel) Insert(ctx context.Context, data *PhaseRelation) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultPhaseRelationModel) FindOne(ctx context.Context, id string) (*PhaseRelation, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data PhaseRelation

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPhaseRelationModel) Update(ctx context.Context, data *PhaseRelation) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultPhaseRelationModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
