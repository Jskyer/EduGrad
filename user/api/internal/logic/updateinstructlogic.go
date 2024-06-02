package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInstructLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInstructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInstructLogic {
	return &UpdateInstructLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInstructLogic) UpdateInstruct(req *types.UpdateInstructReq) (*types.UpdateInstructResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.UpdateInstruct(l.ctx, &user.UpdateInstructReq{
		Id:     req.Id,
		Permit: req.Permit,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UpdateInstructResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
