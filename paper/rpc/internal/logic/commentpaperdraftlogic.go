package logic

import (
	"context"

	"edu-grad/paper/model"
	"edu-grad/paper/rpc/internal/svc"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentPaperDraftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperDraftLogic {
	return &CommentPaperDraftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentPaperDraftLogic) CommentPaperDraft(in *paper.CommentPaperDraftReq) (*paper.CommentPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updateResult, err := l.svcCtx.PaperDraftModel.Update(l.ctx, &model.PaperDraft{
		ID:      oid,
		Comment: in.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &paper.CommentPaperDraftResp{
		Ok: updateResult.ModifiedCount,
	}, nil
}
