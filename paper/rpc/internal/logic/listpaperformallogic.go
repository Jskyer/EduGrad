package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPaperFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperFormalLogic {
	return &ListPaperFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc downloadPaperFormal(DownloadFileReq) returns (DownloadPaperFormalResp);
func (l *ListPaperFormalLogic) ListPaperFormal(in *paper.ListPaperFormalReq) (*paper.ListPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	paperformalMods, err := l.svcCtx.PaperFormalModel.ListByInstructid(l.ctx, in.InstructId)
	if err != nil {
		return nil, err
	}

	paperFormals := make([]*paper.PaperFormal, len(paperformalMods))
	for i, paperformalMod := range paperformalMods {
		paperFormals[i] = common.ConvertPaperFormal(paperformalMod)
	}

	return &paper.ListPaperFormalResp{
		PaperFormals: paperFormals,
	}, nil
}
