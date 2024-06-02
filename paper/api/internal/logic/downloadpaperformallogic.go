package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadPaperFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadPaperFormalLogic {
	return &DownloadPaperFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadPaperFormalLogic) DownloadPaperFormal(req *types.DownloadFileReq) (resp *types.DownloadPaperFormalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
