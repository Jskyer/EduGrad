package logic

import (
	"context"
	"fmt"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	excludeModels, err := l.svcCtx.PhaseRelationModel.ListAllUserids(l.ctx)

	if err != nil {
		return nil, err
	}

	userids := make([]primitive.ObjectID, len(excludeModels))
	for i, datum := range excludeModels {
		oid, err := primitive.ObjectIDFromHex(datum.UserId)

		if err != nil {
			return nil, err
		}

		userids[i] = oid
	}

	fmt.Println(len(userids))

	userModels, totalPage, cnt, err := l.svcCtx.UserModel.FindByPageExceptIds(l.ctx, in.PageNum, in.PageSize, userids)
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
		Count: cnt,
	}, nil
}
