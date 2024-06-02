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

type ListProposalDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProposalDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProposalDraftLogic {
	return &ListProposalDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProposalDraftLogic) ListProposalDraft(req *types.ListProposalDraftReq) (*types.ListProposalDraftResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.InstructId)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.ListProposalDraft(l.ctx, &paper.ListProposalDraftReq{
		InstructId: req.InstructId,
	})

	if err != nil {
		return nil, err
	}

	var resp types.ListProposalDraftResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
