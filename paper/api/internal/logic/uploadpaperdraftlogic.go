package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperDraftLogic {
	return &UploadPaperDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPaperDraftLogic) UploadPaperDraft(req *types.UploadFileReq) (*types.UploadPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.PaperRpc.UploadPaperDraft(l.ctx, &paper.UploadFileReq{
		InstructId: req.InstructId,
		Content:    req.Content,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UploadPaperDraftResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
