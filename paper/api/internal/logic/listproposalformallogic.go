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

type ListProposalFormalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProposalFormalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProposalFormalLogic {
	return &ListProposalFormalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProposalFormalLogic) ListProposalFormal(req *types.ListProposalFormalReq) (*types.ListProposalFormalResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.InstructId)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.ListProposalFormal(l.ctx, &paper.ListProposalFormalReq{
		InstructId: req.InstructId,
	})

	if err != nil {
		return nil, err
	}

	var resp types.ListProposalFormalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
