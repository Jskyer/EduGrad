package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ InstructRelationModel = (*customInstructRelationModel)(nil)

type (
	// InstructRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInstructRelationModel.
	InstructRelationModel interface {
		instructRelationModel
		FindPermitedByStuId(ctx context.Context, stuId string) (*InstructRelation, error)
		FindOneByStuIdTeacherId(ctx context.Context, stuId string, teacherId string) (*InstructRelation, error)
		ListByTeacherId(ctx context.Context, teacherId string) ([]*InstructRelation, error)
		ListByStuId(ctx context.Context, stuId string) ([]*InstructRelation, error)
		CountPermitedByTeacherId(ctx context.Context, teacherId string) (int64, error)
		ListPassedByStuids(ctx context.Context, userids []string) ([]*InstructRelation, error)
	}

	customInstructRelationModel struct {
		*defaultInstructRelationModel
	}
)

// NewInstructRelationModel returns a model for the mongo.
func NewInstructRelationModel(url, db, collection string) InstructRelationModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customInstructRelationModel{
		defaultInstructRelationModel: newDefaultInstructRelationModel(conn),
	}
}

func (m *customInstructRelationModel) FindPermitedByStuId(ctx context.Context, stuId string) (*InstructRelation, error) {
	_, err := primitive.ObjectIDFromHex(stuId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data InstructRelation

	err = m.conn.FindOne(ctx, &data, bson.M{"stuId": stuId, "permit": "1"})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customInstructRelationModel) FindOneByStuIdTeacherId(ctx context.Context, stuId string, teacherId string) (*InstructRelation, error) {
	_, err := primitive.ObjectIDFromHex(stuId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	_, err = primitive.ObjectIDFromHex(teacherId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data InstructRelation

	err = m.conn.FindOne(ctx, &data, bson.M{"stuId": stuId, "teacherId": teacherId})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customInstructRelationModel) ListByTeacherId(ctx context.Context, teacherId string) ([]*InstructRelation, error) {
	_, err := primitive.ObjectIDFromHex(teacherId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data []*InstructRelation

	err = m.conn.Find(ctx, &data, bson.M{"teacherId": teacherId})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *customInstructRelationModel) ListByStuId(ctx context.Context, stuId string) ([]*InstructRelation, error) {
	_, err := primitive.ObjectIDFromHex(stuId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data []*InstructRelation

	err = m.conn.Find(ctx, &data, bson.M{"stuId": stuId})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *customInstructRelationModel) CountPermitedByTeacherId(ctx context.Context, teacherId string) (int64, error) {
	_, err := primitive.ObjectIDFromHex(teacherId)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	cnt, err := m.conn.CountDocuments(ctx, bson.M{"teacherId": teacherId, "permit": "1"})

	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (m *customInstructRelationModel) ListPassedByStuids(ctx context.Context, stuIds []string) ([]*InstructRelation, error) {
	var data []*InstructRelation

	filter := bson.M{
		"permit": "1",
		"state":  "6",
		"stuId":  bson.M{"$in": stuIds},
	}

	err := m.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}
