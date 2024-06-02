package svc

import (
	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	PaperFormalModel    model.PaperFormalModel
	ProposalDraftModel  model.ProposalDraftModel
	ProposalFormalModel model.ProposalFormalModel
	PaperDraftModel     model.PaperDraftModel
	PaperFinalModel     model.PaperFinalModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		PaperFormalModel:    model.NewPaperFormalModel(c.Mongo.URL, c.Mongo.DB, "paperformal"),
		ProposalDraftModel:  model.NewProposalDraftModel(c.Mongo.URL, c.Mongo.DB, "proposaldraft"),
		ProposalFormalModel: model.NewProposalFormalModel(c.Mongo.URL, c.Mongo.DB, "proposalformal"),
		PaperDraftModel:     model.NewPaperDraftModel(c.Mongo.URL, c.Mongo.DB, "paperdraft"),
		PaperFinalModel:     model.NewPaperFinalModel(c.Mongo.URL, c.Mongo.DB, "paperfinal"),
	}
}
