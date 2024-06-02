package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStateLogic {
	return &UpdateStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStateLogic) UpdateState(req *types.UpdateStateReq) (*types.UpdateStateResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.UpdateState(l.ctx, &user.UpdateStateReq{
		Id:    req.Id,
		State: req.State,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UpdateStateResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
