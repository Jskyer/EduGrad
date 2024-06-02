package logic

import (
	"context"
	"errors"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProcessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcessLogic {
	return &UpdateProcessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProcessLogic) UpdateProcess(req *types.UpdateProcessReq) (*types.UpdateProcessResp, error) {
	// todo: add your logic here and delete this line
	if req.Process != 2 {
		return nil, errors.New("process不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.UpdateProcess(l.ctx, &user.UpdateProcessReq{
		Id:      req.Id,
		Process: req.Process,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UpdateProcessResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
