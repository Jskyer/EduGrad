package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPhaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhaseLogic {
	return &AddPhaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPhaseLogic) AddPhase(req *types.AddPhaseReq) (*types.AddPhaseResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.AddPhase(l.ctx, &user.AddPhaseReq{
		Term:    req.Term,
		Process: req.Process,
	})

	if err != nil {
		return nil, err
	}

	var resp types.AddPhaseResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
