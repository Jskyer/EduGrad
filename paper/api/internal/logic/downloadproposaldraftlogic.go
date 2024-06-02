package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadProposalDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadProposalDraftLogic {
	return &DownloadProposalDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadProposalDraftLogic) DownloadProposalDraft(req *types.DownloadFileReq) (resp *types.DownloadProposalDraftResp, err error) {
	// todo: add your logic here and delete this line

	return
}
