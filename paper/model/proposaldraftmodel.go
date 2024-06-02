package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ ProposalDraftModel = (*customProposalDraftModel)(nil)

type (
	// ProposalDraftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProposalDraftModel.
	ProposalDraftModel interface {
		proposalDraftModel
		ListByInstructid(ctx context.Context, instructId string) ([]*ProposalDraft, error)
	}

	customProposalDraftModel struct {
		*defaultProposalDraftModel
	}
)

// NewProposalDraftModel returns a model for the mongo.
func NewProposalDraftModel(url, db, collection string) ProposalDraftModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProposalDraftModel{
		defaultProposalDraftModel: newDefaultProposalDraftModel(conn),
	}
}

func (m *customProposalDraftModel) ListByInstructid(ctx context.Context, instructId string) ([]*ProposalDraft, error) {
	var data []*ProposalDraft

	err := m.conn.Find(ctx, &data, bson.M{"instructId": instructId})

	if err != nil {
		return nil, err
	}

	return data, nil
}
