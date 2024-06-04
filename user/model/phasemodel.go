package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ PhaseModel = (*customPhaseModel)(nil)

type (
	// PhaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPhaseModel.
	PhaseModel interface {
		phaseModel
		ListByTermSorted(ctx context.Context, pageNum int64, pageSize int64) ([]*Phase, int64, int64, error)
		FindCurrentPhase(ctx context.Context) (*Phase, error)
		FindOneByTerm(ctx context.Context, term string) (*Phase, error)
	}

	customPhaseModel struct {
		*defaultPhaseModel
	}
)

// NewPhaseModel returns a model for the mongo.
func NewPhaseModel(url, db, collection string) PhaseModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPhaseModel{
		defaultPhaseModel: newDefaultPhaseModel(conn),
	}
}

func (m *customPhaseModel) ListByTermSorted(ctx context.Context, pageNum int64, pageSize int64) ([]*Phase, int64, int64, error) {
	var data []*Phase

	opts := new(options.FindOptions)
	opts.SetSkip((pageNum - 1) * pageSize)
	opts.SetLimit(pageSize)

	sort := bson.D{
		bson.E{Key: "term", Value: -1},
	}
	opts.SetSort(sort)

	err := m.conn.Find(ctx, &data, bson.M{}, opts)
	if err != nil {
		return nil, 0, 0, err
	}

	// 对应filter的记录总数
	cnt, err := m.conn.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, 0, err
	}

	// 转换为总页数
	totalPage := (cnt + pageSize - 1) / pageSize

	return data, totalPage, cnt, nil
}

func (m *customPhaseModel) FindCurrentPhase(ctx context.Context) (*Phase, error) {
	var data Phase

	err := m.conn.FindOne(ctx, &data, bson.M{"process": 2})
	if err != nil {
		return nil, err
	}

	return &data, err
}

func (m *customPhaseModel) FindOneByTerm(ctx context.Context, term string) (*Phase, error) {
	var data Phase

	err := m.conn.FindOne(ctx, &data, bson.M{"term": term})
	if err != nil {
		return nil, err
	}

	return &data, err
}
