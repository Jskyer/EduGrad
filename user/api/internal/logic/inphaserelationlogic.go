package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InphaseRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInphaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InphaseRelationLogic {
	return &InphaseRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InphaseRelationLogic) InphaseRelation(req *types.InphaseRelationReq) (*types.InphaseRelationResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Userid)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.UserRpc.InphaseRelation(l.ctx, &user.InphaseRelationReq{
		Userid: req.Userid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.InphaseRelationResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
