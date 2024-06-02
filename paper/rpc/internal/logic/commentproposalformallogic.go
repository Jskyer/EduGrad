package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentProposalFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentProposalFormalLogic {
	return &CommentProposalFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentProposalFormalLogic) CommentProposalFormal(in *paper.CommentProposalFormalReq) (*paper.CommentProposalFormalResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.ProposalFormalModel.Update(l.ctx, &model.ProposalFormal{
		ID:      oid,
		Comment: in.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CommentProposalFormalResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
