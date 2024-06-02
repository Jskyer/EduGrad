package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPaperFinalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadPaperFinalLogic {
	return &DownloadPaperFinalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadPaperFinalLogic) DownloadPaperFinal(req *types.DownloadFileReq) (resp *types.DownloadPaperFinalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
