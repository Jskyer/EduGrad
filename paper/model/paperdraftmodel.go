package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ PaperDraftModel = (*customPaperDraftModel)(nil)

type (
	// PaperDraftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaperDraftModel.
	PaperDraftModel interface {
		paperDraftModel
		ListByInstructid(ctx context.Context, instructId string) ([]*PaperDraft, error)
	}

	customPaperDraftModel struct {
		*defaultPaperDraftModel
	}
)

// NewPaperDraftModel returns a model for the mongo.
func NewPaperDraftModel(url, db, collection string) PaperDraftModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPaperDraftModel{
		defaultPaperDraftModel: newDefaultPaperDraftModel(conn),
	}
}

func (m *customPaperDraftModel) ListByInstructid(ctx context.Context, instructId string) ([]*PaperDraft, error) {
	var data []*PaperDraft

	err := m.conn.Find(ctx, &data, bson.M{"instructId": instructId})

	if err != nil {
		return nil, err
	}

	return data, nil
}
