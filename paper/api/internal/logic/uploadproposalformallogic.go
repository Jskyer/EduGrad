package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadProposalFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadProposalFormalLogic {
	return &UploadProposalFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadProposalFormalLogic) UploadProposalFormal(req *types.UploadFileReq) (*types.UploadProposalFormalResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.PaperRpc.UploadProposalFormal(l.ctx, &paper.UploadFileReq{
		InstructId: req.InstructId,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UploadProposalFormalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
