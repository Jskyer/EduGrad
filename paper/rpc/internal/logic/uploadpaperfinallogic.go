package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperFinalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperFinalLogic {
	return &UploadPaperFinalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPaperFinalLogic) UploadPaperFinal(in *paper.UploadFileReq) (*paper.UploadPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PaperFinalModel.Insert(l.ctx, &model.PaperFinal{
		InstructId: in.InstructId,
		Content:    in.Content,
	})

	if err != nil {
		return nil, err
	}

	return &paper.UploadPaperFinalResp{
		Filename: in.Content,
	}, nil
}
