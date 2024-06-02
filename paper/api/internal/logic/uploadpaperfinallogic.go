package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperFinalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperFinalLogic {
	return &UploadPaperFinalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPaperFinalLogic) UploadPaperFinal(req *types.UploadFileReq) (*types.UploadPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.PaperRpc.UploadPaperFinal(l.ctx, &paper.UploadFileReq{
		InstructId: req.InstructId,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UploadPaperFinalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
