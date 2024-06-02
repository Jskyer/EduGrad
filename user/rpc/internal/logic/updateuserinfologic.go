package logic

import (
	"context"
	"fmt"

	"edu-grad/common"
	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line

	// 检查id合法性
	targetUser := in.User
	oid, err := primitive.ObjectIDFromHex(targetUser.Id)
	if err != nil {
		fmt.Println("ObjectIDFromHex err:", err)
		return nil, err
	}

	// 检查是否存在对应id user，不存在直接返回，mongodb不需要检查，id不存在不会添加或更新
	// existUser, err := l.svcCtx.UserModel.FindOne(l.ctx, targetUser.Id)
	// if err != nil {
	// 	fmt.Println("FindOne err:", err)
	// 	return nil, err
	// }
	// if existUser == nil {
	// 	return &user.UpdateUserInfoResp{
	// 		Updated: 0,
	// 	}, nil
	// }

	if targetUser.Password == "" && targetUser.Grade == "" {
		return &user.UpdateUserInfoResp{
			Updated: 0,
		}, nil
	}

	newUser := model.User{}
	newUser.ID = oid

	if targetUser.Password != "" {
		newUser.Password = common.Md5EncodeUpper(targetUser.Password)
	}

	if targetUser.Grade != "" {
		newUser.Grade = targetUser.Grade
	}

	res, err := l.svcCtx.UserModel.Update(l.ctx, &newUser)
	if err != nil {
		fmt.Println("Update err")
		return nil, err
	}
	fmt.Println("res.ModifiedCount:", res.ModifiedCount)

	return &user.UpdateUserInfoResp{
		Updated: res.ModifiedCount,
	}, nil
}
