package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProposalDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProposalDraftLogic {
	return &ListProposalDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProposalDraftLogic) ListProposalDraft(in *paper.ListProposalDraftReq) (*paper.ListProposalDraftResp, error) {
	// todo: add your logic here and delete this line
	proposalDraftMods, err := l.svcCtx.ProposalDraftModel.ListByInstructid(l.ctx, in.InstructId)
	if err != nil {
		return nil, err
	}

	proposalDrafts := make([]*paper.ProposalDraft, len(proposalDraftMods))
	for i, proposalDraftMod := range proposalDraftMods {
		proposalDrafts[i] = common.ConvertProposalDraft(proposalDraftMod)
	}

	return &paper.ListProposalDraftResp{
		ProposalDrafts: proposalDrafts,
	}, nil
}
