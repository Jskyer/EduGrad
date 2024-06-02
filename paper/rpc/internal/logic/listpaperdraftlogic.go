package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPaperDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperDraftLogic {
	return &ListPaperDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPaperDraftLogic) ListPaperDraft(in *paper.ListPaperDraftReq) (*paper.ListPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	paperdraftMods, err := l.svcCtx.PaperDraftModel.ListByInstructid(l.ctx, in.InstructId)
	if err != nil {
		return nil, err
	}

	paperDrafts := make([]*paper.PaperDraft, len(paperdraftMods))
	for i, paperdraftMod := range paperdraftMods {
		paperDrafts[i] = common.ConvertPaperDraft(paperdraftMod)
	}

	return &paper.ListPaperDraftResp{
		PaperDrafts: paperDrafts,
	}, nil
}
