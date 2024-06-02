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

type ListPaperFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPaperFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperFormalLogic {
	return &ListPaperFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPaperFormalLogic) ListPaperFormal(req *types.ListPaperFormalReq) (*types.ListPaperFormalResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.InstructId)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.ListPaperFormal(l.ctx, &paper.ListPaperFormalReq{
		InstructId: req.InstructId,
	})

	if err != nil {
		return nil, err
	}

	var resp types.ListPaperFormalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
