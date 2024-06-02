package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadProposalDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadProposalDraftLogic {
	return &UploadProposalDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadProposalDraftLogic) UploadProposalDraft(req *types.UploadFileReq) (*types.UploadProposalDraftResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.PaperRpc.UploadProposalDraft(l.ctx, &paper.UploadFileReq{
		InstructId: req.InstructId,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UploadProposalDraftResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
