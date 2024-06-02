package logic

import (
	"context"
	"fmt"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (*types.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.UpdateUserInfoReq{
		User: &user.UserInfo{
			Id:       req.ID,
			Password: req.Password,
			Grade:    req.Grade,
		},
	})

	if err != nil {
		return nil, err
	}

	var resp types.UpdateUserInfoResp
	_ = copier.Copy(&resp, rpcResp)
	fmt.Printf("resp: %+v\n", resp)

	return &resp, nil
}
