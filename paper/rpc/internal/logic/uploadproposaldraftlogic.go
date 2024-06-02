package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadProposalDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadProposalDraftLogic {
	return &UploadProposalDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadProposalDraftLogic) UploadProposalDraft(in *paper.UploadFileReq) (*paper.UploadProposalDraftResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ProposalDraftModel.Insert(l.ctx, &model.ProposalDraft{
		InstructId: in.InstructId,
		Content:    in.Content,
	})

	if err != nil {
		return nil, err
	}

	return &paper.UploadProposalDraftResp{
		Filename: in.Content,
	}, nil
}
