package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ PhaseRelationModel = (*customPhaseRelationModel)(nil)

type (
	// PhaseRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhaseRelationModel.
	PhaseRelationModel interface {
		phaseRelationModel
		InsertMany(ctx context.Context, data []*PhaseRelation) error
		ExistUseridsInPhaseRelation(ctx context.Context, phaseId string, userids []string) ([]*PhaseRelation, error)
		FindOneByPhaseid(ctx context.Context, phaseId string) (*PhaseRelation, error)
		FindOneByUserid(ctx context.Context, userid string) (*PhaseRelation, error)
		ListByPhaseidAndIdentity(ctx context.Context, phaseId string, identity string, pageNum int64, pageSize int64) ([]*PhaseRelation, int64, int64, error)
		ListAllByPhaseidAndIdentity(ctx context.Context, phaseId string, identity string) ([]*PhaseRelation, error)
		UpdateStudentsPass(ctx context.Context, phaseId string, userids []string, pass int64) (*mongo.UpdateResult, error)
		ListTeachersByPhaseidExceptTids(ctx context.Context, phaseId string, userids []string, pageNum int64, pageSize int64) ([]*PhaseRelation, int64, int64, error)
		ListAllUserids(ctx context.Context) ([]*PhaseRelation, error)
	}

	customPhaseRelationModel struct {
		*defaultPhaseRelationModel
	}
)

// NewPhaseRelationModel returns a model for the mongo.
func NewPhaseRelationModel(url, db, collection string) PhaseRelationModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPhaseRelationModel{
		defaultPhaseRelationModel: newDefaultPhaseRelationModel(conn),
	}
}

func (m *customPhaseRelationModel) InsertMany(ctx context.Context, data []*PhaseRelation) error {
	insertData := make([]interface{}, len(data))

	for i, datum := range data {
		if datum.ID.IsZero() {
			datum.ID = primitive.NewObjectID()
			datum.CreateAt = time.Now()
			datum.UpdateAt = time.Now()
		}

		insertData[i] = datum
	}

	_, err := m.conn.InsertMany(ctx, insertData)
	if err != nil {
		return err
	}

	return nil
}

func (m *customPhaseRelationModel) ExistUseridsInPhaseRelation(ctx context.Context, phaseId string, userids []string) ([]*PhaseRelation, error) {
	var data []*PhaseRelation

	filter := bson.M{
		"phaseId": phaseId,
		"userId": bson.M{
			"$in": userids,
		},
	}

	err := m.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *customPhaseRelationModel) FindOneByPhaseid(ctx context.Context, phaseId string) (*PhaseRelation, error) {
	var data PhaseRelation

	filter := bson.M{
		"phaseId": phaseId,
	}

	err := m.conn.FindOne(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *customPhaseRelationModel) ListByPhaseidAndIdentity(ctx context.Context, phaseId string, identity string, pageNum int64, pageSize int64) ([]*PhaseRelation, int64, int64, error) {
	var data []*PhaseRelation

	filter := bson.M{"phaseId": phaseId, "identity": identity}

	opts := new(options.FindOptions)
	opts.SetSkip((pageNum - 1) * pageSize)
	opts.SetLimit(pageSize)

	err := m.conn.Find(ctx, &data, filter, opts)
	if err != nil {
		return nil, 0, 0, err
	}

	cnt, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	// 转换为总页数
	totalPage := (cnt + pageSize - 1) / pageSize

	return data, totalPage, cnt, nil
}

func (m *customPhaseRelationModel) ListAllByPhaseidAndIdentity(ctx context.Context, phaseId string, identity string) ([]*PhaseRelation, error) {
	var data []*PhaseRelation

	filter := bson.M{"phaseId": phaseId, "identity": identity}

	err := m.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *customPhaseRelationModel) UpdateStudentsPass(ctx context.Context, phaseId string, userids []string, pass int64) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"phaseId":  phaseId,
		"identity": "学生",
		"userId": bson.M{
			"$in": userids,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"pass": pass,
		},
	}

	updateResult, err := m.conn.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (m *customPhaseRelationModel) ListTeachersByPhaseidExceptTids(ctx context.Context, phaseId string, userids []string, pageNum int64, pageSize int64) ([]*PhaseRelation, int64, int64, error) {
	var data []*PhaseRelation

	filter := bson.M{
		"phaseId":  phaseId,
		"identity": "导师",
		"userId": bson.M{
			"$nin": userids,
		},
	}

	opts := new(options.FindOptions)
	opts.SetSkip((pageNum - 1) * pageSize)
	opts.SetLimit(pageSize)

	err := m.conn.Find(ctx, &data, filter, opts)
	if err != nil {
		return nil, 0, 0, err
	}

	cnt, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	// 转换为总页数
	totalPage := (cnt + pageSize - 1) / pageSize

	return data, totalPage, cnt, nil
}

func (m *customPhaseRelationModel) FindOneByUserid(ctx context.Context, userid string) (*PhaseRelation, error) {
	var data PhaseRelation

	filter := bson.M{
		"userId": userid,
	}

	err := m.conn.FindOne(ctx, &data, filter)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *customPhaseRelationModel) ListAllUserids(ctx context.Context) ([]*PhaseRelation, error) {
	var data []*PhaseRelation

	project := bson.M{
		"userId": 1,
		"_id":    0,
	}
	opts := options.FindOptions{Projection: project}

	// opts := new(options.FindOptions)
	// opts.SetProjection(project)

	err := m.conn.Find(ctx, &data, bson.M{}, &opts)
	if err != nil {
		return nil, err
	}

	return data, nil
}
