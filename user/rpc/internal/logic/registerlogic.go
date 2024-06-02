package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/common"
	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	if data != nil {
		return nil, errors.New("用户已存在")
	}

	if err != mon.ErrNotFound {
		return nil, err
	}

	fmt.Println(err)

	// 加密
	encoded := common.Md5EncodeUpper(in.Password)

	err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: encoded,
		Identity: in.Identity,
		Grade:    in.Grade,
	})

	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		Ok: 1,
	}, nil
}
