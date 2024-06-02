package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPaperFinalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperFinalLogic {
	return &ListPaperFinalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPaperFinalLogic) ListPaperFinal(in *paper.ListPaperFinalReq) (*paper.ListPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	paperfinalMods, err := l.svcCtx.PaperFinalModel.ListByInstructid(l.ctx, in.InstructId)
	if err != nil {
		return nil, err
	}

	paperFinals := make([]*paper.PaperFinal, len(paperfinalMods))
	for i, paperfinalMod := range paperfinalMods {
		paperFinals[i] = common.ConvertPaperFinal(paperfinalMod)
	}

	return &paper.ListPaperFinalResp{
		PaperFinals: paperFinals,
	}, nil
}
