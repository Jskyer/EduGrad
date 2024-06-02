package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentPaperFormalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperFormalLogic {
	return &CommentPaperFormalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentPaperFormalLogic) CommentPaperFormal(in *paper.CommentPaperFormalReq) (*paper.CommentPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.PaperFormalModel.Update(l.ctx, &model.PaperFormal{
		ID:      oid,
		Comment: in.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CommentPaperFormalResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
