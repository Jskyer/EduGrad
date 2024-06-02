package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadProposalFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadProposalFormalLogic {
	return &DownloadProposalFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadProposalFormalLogic) DownloadProposalFormal(req *types.DownloadFileReq) (resp *types.DownloadProposalFormalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
