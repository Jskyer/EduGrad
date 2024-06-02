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

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line

	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Identity: req.Identity,
		Grade:    req.Grade,
	})

	if err != nil {
		return nil, err
	}

	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)

	fmt.Printf("resp: %+v\n", resp)
	return &resp, nil
}
