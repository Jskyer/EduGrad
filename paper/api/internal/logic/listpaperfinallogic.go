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

type ListPaperFinalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperFinalLogic {
	return &ListPaperFinalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPaperFinalLogic) ListPaperFinal(req *types.ListPaperFinalReq) (*types.ListPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.InstructId)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.ListPaperFinal(l.ctx, &paper.ListPaperFinalReq{
		InstructId: req.InstructId,
	})

	if err != nil {
		return nil, err
	}

	var resp types.ListPaperFinalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
