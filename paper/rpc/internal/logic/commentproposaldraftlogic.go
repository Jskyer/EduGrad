package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentProposalDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentProposalDraftLogic {
	return &CommentProposalDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentProposalDraftLogic) CommentProposalDraft(in *paper.CommentProposalDraftReq) (*paper.CommentProposalDraftResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.ProposalDraftModel.Update(l.ctx, &model.ProposalDraft{
		ID:      oid,
		Comment: in.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CommentProposalDraftResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
