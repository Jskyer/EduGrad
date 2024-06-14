package logic

import (
	"context"
	"fmt"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userModels, totalPage, cnt, err := l.svcCtx.UserModel.FindPageByGradeAndIdentity(l.ctx, in.Identity, in.Grade, in.PageNum, in.PageSize, userids)
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
		Count: cnt,
	}, nil
}
