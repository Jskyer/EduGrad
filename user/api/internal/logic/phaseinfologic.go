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

type PhaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhaseInfoLogic {
	return &PhaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhaseInfoLogic) PhaseInfo(req *types.PhaseInfoReq) (*types.PhaseInfoResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Phaseid)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.UserRpc.GetPhaseinfo(l.ctx, &user.PhaseInfoReq{
		Phaseid: req.Phaseid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PhaseInfoResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
