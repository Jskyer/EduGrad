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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("enter api login logic")

	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	//
	// fmt.Printf("loginResp: %+v\n", loginResp)

	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)

	fmt.Printf("resp: %+v\n", resp)

	//
	// resp.Id = "111"
	// resp.Name = "jack"
	// resp.Token = "gg"

	// fmt.Println("=======")

	// fmt.Printf("loginResp: %+v\n", loginResp)
	// fmt.Printf("resp: %+v\n", resp)

	return &resp, nil
}
