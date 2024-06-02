package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ ProposalFormalModel = (*customProposalFormalModel)(nil)

type (
	// ProposalFormalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProposalFormalModel.
	ProposalFormalModel interface {
		proposalFormalModel
		ListByInstructid(ctx context.Context, instructId string) ([]*ProposalFormal, error)
	}

	customProposalFormalModel struct {
		*defaultProposalFormalModel
	}
)

// NewProposalFormalModel returns a model for the mongo.
func NewProposalFormalModel(url, db, collection string) ProposalFormalModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customProposalFormalModel{
		defaultProposalFormalModel: newDefaultProposalFormalModel(conn),
	}
}

func (m *customProposalFormalModel) ListByInstructid(ctx context.Context, instructId string) ([]*ProposalFormal, error) {
	var data []*ProposalFormal

	err := m.conn.Find(ctx, &data, bson.M{"instructId": instructId})

	if err != nil {
		return nil, err
	}

	return data, nil
}
