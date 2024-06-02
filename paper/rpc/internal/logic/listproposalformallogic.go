package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProposalFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProposalFormalLogic {
	return &ListProposalFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProposalFormalLogic) ListProposalFormal(in *paper.ListProposalFormalReq) (*paper.ListProposalFormalResp, error) {
	// todo: add your logic here and delete this line
	proposalFormalMods, err := l.svcCtx.ProposalFormalModel.ListByInstructid(l.ctx, in.InstructId)
	if err != nil {
		return nil, err
	}

	proposalFormals := make([]*paper.ProposalFormal, len(proposalFormalMods))
	for i, proposalFormalMod := range proposalFormalMods {
		proposalFormals[i] = common.ConvertProposalFormal(proposalFormalMod)
	}

	return &paper.ListProposalFormalResp{
		ProposalFormals: proposalFormals,
	}, nil
}
