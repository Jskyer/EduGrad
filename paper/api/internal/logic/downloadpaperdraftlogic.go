package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPaperDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadPaperDraftLogic {
	return &DownloadPaperDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadPaperDraftLogic) DownloadPaperDraft(req *types.DownloadFileReq) (resp *types.DownloadPaperDraftResp, err error) {
	// todo: add your logic here and delete this line

	return
}
