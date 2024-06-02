package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddPhaseRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhaseRelationLogic {
	return &AddPhaseRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPhaseRelationLogic) AddPhaseRelation(req *types.AddPhaseRelationReq) (*types.AddPhaseRelationResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("id不合法")
	}

	if len(req.Users) <= 0 && req.Grade == "" {
		fmt.Println("缺少添加的用户信息")
		return nil, errors.New("缺少添加的用户信息")
	}

	if len(req.Users) > 0 {
		for _, userId := range req.Users {
			_, err := primitive.ObjectIDFromHex(userId)
			if err != nil {
				return nil, errors.New("users不合法")
			}
		}
	}

	rpcResp, err := l.svcCtx.UserRpc.AddPhaseRelation(l.ctx, &user.AddPhaseRelationReq{
		Id:    req.Id,
		Users: req.Users,
		Grade: req.Grade,
	})

	if err != nil {
		return nil, err
	}

	var resp types.AddPhaseRelationResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
