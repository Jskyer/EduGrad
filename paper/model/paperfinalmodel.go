package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ PaperFinalModel = (*customPaperFinalModel)(nil)

type (
	// PaperFinalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaperFinalModel.
	PaperFinalModel interface {
		paperFinalModel
		ListByInstructid(ctx context.Context, instructId string) ([]*PaperFinal, error)
	}

	customPaperFinalModel struct {
		*defaultPaperFinalModel
	}
)

// NewPaperFinalModel returns a model for the mongo.
func NewPaperFinalModel(url, db, collection string) PaperFinalModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPaperFinalModel{
		defaultPaperFinalModel: newDefaultPaperFinalModel(conn),
	}
}

func (m *customPaperFinalModel) ListByInstructid(ctx context.Context, instructId string) ([]*PaperFinal, error) {
	var data []*PaperFinal

	err := m.conn.Find(ctx, &data, bson.M{"instructId": instructId})

	if err != nil {
		return nil, err
	}

	return data, nil
}
