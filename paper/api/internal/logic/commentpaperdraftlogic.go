package logic

import (
	"context"

	"edu-grad/paper/api/internal/svc"
	"edu-grad/paper/api/internal/types"
	"edu-grad/paper/rpc/types/proto/paper"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentPaperDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperDraftLogic {
	return &CommentPaperDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPaperDraftLogic) CommentPaperDraft(req *types.CommentPaperDraftReq) (*types.CommentPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.CommentPaperDraft(l.ctx, &paper.CommentPaperDraftReq{
		Id:      req.Id,
		Comment: req.Comment,
	})

	if err != nil {
		return nil, err
	}

	var resp types.CommentPaperDraftResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
