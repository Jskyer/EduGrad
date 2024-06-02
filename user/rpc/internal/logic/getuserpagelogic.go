package logic

import (
	"context"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPageLogic) GetUserPage(in *user.GetUserPageReq) (*user.GetUserPageResp, error) {
	// todo: add your logic here and delete this line
	userModels, totalPage, err := l.svcCtx.UserModel.FindByPage(l.ctx, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	var users []*user.UserInfo
	for _, umodel := range userModels {
		users = append(users, &user.UserInfo{
			Id:       umodel.ID.Hex(),
			Username: umodel.Username,
			Identity: umodel.Identity,
			Grade:    umodel.Grade,
		})
	}

	return &user.GetUserPageResp{
		Users: users,
		Total: totalPage,
	}, nil
}
