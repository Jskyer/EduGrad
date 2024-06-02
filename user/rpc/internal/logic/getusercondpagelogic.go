package logic

import (
	"context"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCondPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCondPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCondPageLogic {
	return &GetUserCondPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCondPageLogic) GetUserCondPage(in *user.GetUserCondPageReq) (*user.GetUserCondPageResp, error) {
	// todo: add your logic here and delete this line
	userModels, totalPage, err := l.svcCtx.UserModel.FindPageByGradeAndIdentity(l.ctx, in.Identity, in.Grade, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	users := make([]*user.UserInfo, len(userModels))
	for i, umodel := range userModels {
		users[i] = &user.UserInfo{
			Id:       umodel.ID.Hex(),
			Username: umodel.Username,
			Identity: umodel.Identity,
			Grade:    umodel.Grade,
		}
	}

	return &user.GetUserCondPageResp{
		Users: users,
		Total: totalPage,
	}, nil
}
