package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Id)
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err == mon.ErrNotFound {
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResp{
		User: &user.UserInfo{
			Id:       res.ID.Hex(),
			Username: res.Username,
			Identity: res.Identity,
			Grade:    res.Grade,
		},
	}, nil
}
