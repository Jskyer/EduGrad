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

type CommentPaperFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperFormalLogic {
	return &CommentPaperFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPaperFormalLogic) CommentPaperFormal(req *types.CommentPaperFormalReq) (*types.CommentPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.CommentPaperFormal(l.ctx, &paper.CommentPaperFormalReq{
		Id:      req.Id,
		Comment: req.Comment,
	})

	if err != nil {
		return nil, err
	}

	var resp types.CommentPaperFormalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
