package logic

import (
	"context"
	"errors"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EndPhaseRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndPhaseRelationLogic {
	return &EndPhaseRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EndPhaseRelationLogic) EndPhaseRelation(req *types.EndPhaseRelationReq) (*types.EndPhaseRelationResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("id不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.EndPhaseRelation(l.ctx, &user.EndPhaseRelationReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var resp types.EndPhaseRelationResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
