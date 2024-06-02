package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ PaperFormalModel = (*customPaperFormalModel)(nil)

type (
	// PaperFormalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaperFormalModel.
	PaperFormalModel interface {
		paperFormalModel
		ListByInstructid(ctx context.Context, instructId string) ([]*PaperFormal, error)
	}

	customPaperFormalModel struct {
		*defaultPaperFormalModel
	}
)

// NewPaperFormalModel returns a model for the mongo.
func NewPaperFormalModel(url, db, collection string) PaperFormalModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPaperFormalModel{
		defaultPaperFormalModel: newDefaultPaperFormalModel(conn),
	}
}

func (m *customPaperFormalModel) ListByInstructid(ctx context.Context, instructId string) ([]*PaperFormal, error) {
	var data []*PaperFormal

	err := m.conn.Find(ctx, &data, bson.M{"instructId": instructId})

	if err != nil {
		return nil, err
	}

	return data, nil
}
