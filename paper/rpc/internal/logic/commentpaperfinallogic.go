package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentPaperFinalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperFinalLogic {
	return &CommentPaperFinalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentPaperFinalLogic) CommentPaperFinal(in *paper.CommentPaperFinalReq) (*paper.CommentPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.PaperFinalModel.Update(l.ctx, &model.PaperFinal{
		ID:      oid,
		Comment: in.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CommentPaperFinalResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
