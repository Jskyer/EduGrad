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

type ListPaperDraftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPaperDraftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaperDraftLogic {
	return &ListPaperDraftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPaperDraftLogic) ListPaperDraft(req *types.ListPaperDraftReq) (*types.ListPaperDraftResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.InstructId)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.ListPaperDraft(l.ctx, &paper.ListPaperDraftReq{
		InstructId: req.InstructId,
	})

	if err != nil {
		return nil, err
	}

	var resp types.ListPaperDraftResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
