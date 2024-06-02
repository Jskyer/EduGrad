package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTitleLogic {
	return &UpdateTitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTitleLogic) UpdateTitle(req *types.UpdateTitleReq) (*types.UpdateTitleResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.UpdateTitle(l.ctx, &user.UpdateTitleReq{
		Id:    req.Id,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UpdateTitleResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
