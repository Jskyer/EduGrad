package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadProposalFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadProposalFormalLogic {
	return &UploadProposalFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadProposalFormalLogic) UploadProposalFormal(in *paper.UploadFileReq) (*paper.UploadProposalFormalResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ProposalFormalModel.Insert(l.ctx, &model.ProposalFormal{
		InstructId: in.InstructId,
		Content:    in.Content,
	})

	if err != nil {
		return nil, err
	}

	return &paper.UploadProposalFormalResp{
		Filename: in.Content,
	}, nil
}
