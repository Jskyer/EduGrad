package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	logx.Info("enter rpc login logic")

	data, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	if data == nil {
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		return nil, err
	}

	logx.Info("finish find")

	// 验证密码
	valid := common.Md5ValidateUpper(in.Password, data.Password)
	if !valid {
		return nil, errors.New("密码不正确")
	}

	// 生成token
	tokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := tokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: data.ID.Hex(),
	})

	fmt.Println("tokenResp: ", tokenResp)

	if err != nil {
		return nil, err
	}

	expireAt := time.Unix(tokenResp.AccessExpire, 0).String()

	expireTime := (int)(l.svcCtx.Config.JwtAuth.AccessExpire)
	// logx.Info(expireTime)

	// 保存登录用户到redis，方便之后查询完整用户数据
	rkey := "user:" + data.ID.Hex()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// err = l.svcCtx.RdsCli.Set(rkey, string(jsonData))
	err = l.svcCtx.RdsCli.Setex(rkey, string(jsonData), expireTime)
	if err != nil {
		return nil, err
	}

	logx.Info("save user to redis")

	return &user.LoginResp{
		Id:       data.ID.Hex(),
		Name:     data.Username,
		Token:    tokenResp.AccessToken,
		ExpireAt: expireAt,
	}, nil
}
