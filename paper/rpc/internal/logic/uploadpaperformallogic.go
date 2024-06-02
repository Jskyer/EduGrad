package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPaperFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPaperFormalLogic {
	return &UploadPaperFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPaperFormalLogic) UploadPaperFormal(in *paper.UploadFileReq) (*paper.UploadPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.PaperFormalModel.Insert(l.ctx, &model.PaperFormal{
		InstructId: in.InstructId,
		Content:    in.Content,
	})

	if err != nil {
		return nil, err
	}

	return &paper.UploadPaperFormalResp{
		Filename: in.Content,
	}, nil
}
