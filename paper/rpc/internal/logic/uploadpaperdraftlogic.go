package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperDraftLogic {
	return &UploadPaperDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPaperDraftLogic) UploadPaperDraft(in *paper.UploadFileReq) (*paper.UploadPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PaperDraftModel.Insert(l.ctx, &model.PaperDraft{
		InstructId: in.InstructId,
		Content:    in.Content,
	})

	if err != nil {
		return nil, err
	}

	return &paper.UploadPaperDraftResp{
		Filename: in.Content,
	}, nil
}
