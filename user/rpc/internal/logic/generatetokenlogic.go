package logic

import (
	"context"
	"fmt"
	"time"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *user.GenerateTokenReq) (*user.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	token, err := common.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, expire, in.UserId)

	if err != nil {
		fmt.Println("GenerateToken err: ", err)
		return nil, err
	}

	return &user.GenerateTokenResp{
		AccessToken:  token,
		AccessExpire: now + expire, //到期时间
	}, nil
}
