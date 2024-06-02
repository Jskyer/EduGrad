package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperFormalLogic {
	return &UploadPaperFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPaperFormalLogic) UploadPaperFormal(req *types.UploadFileReq) (*types.UploadPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.PaperRpc.UploadPaperFormal(l.ctx, &paper.UploadFileReq{
		InstructId: req.InstructId,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UploadPaperFormalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
