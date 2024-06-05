package logic

import (
	"context"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type InphaseRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInphaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InphaseRelationLogic {
	return &InphaseRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InphaseRelationLogic) InphaseRelation(in *user.InphaseRelationReq) (*user.InphaseRelationResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.PhaseRelationModel.FindOneByUserid(l.ctx, in.Userid)

	if err != nil && err != mon.ErrNotFound {
		return nil, err
	}

	if err == mon.ErrNotFound {
		return &user.InphaseRelationResp{Phaseid: ""}, nil
	}

	return &user.InphaseRelationResp{
		Phaseid: data.PhaseId,
	}, nil
}
