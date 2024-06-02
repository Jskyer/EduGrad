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

type CheckPaperFinalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckPaperFinalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPaperFinalLogic {
	return &CheckPaperFinalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckPaperFinalLogic) CheckPaperFinal(req *types.CheckPaperFinalReq) (*types.CheckPaperFinalResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.PaperRpc.CheckPaperFinal(l.ctx, &paper.CheckPaperFinalReq{
		Id:          req.Id,
		CheckResult: req.CheckResult,
	})

	if err != nil {
		return nil, err
	}

	var resp types.CheckPaperFinalResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
