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

type PhaseCommittedtistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhaseCommittedtistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhaseCommittedtistLogic {
	return &PhaseCommittedtistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhaseCommittedtistLogic) PhaseCommittedtist(req *types.PhaseCommittedtistReq) (*types.PhaseCommittedtistResp, error) {
	// todo: add your logic here and delete this line
	if req.PageNum <= 0 || req.PageSize <= 0 {
		return nil, errors.New("pageNum或pageSize不合法")
	}

	_, err := primitive.ObjectIDFromHex(req.Stuid)
	if err != nil {
		return nil, err
	}

	rpcResp, err := l.svcCtx.UserRpc.PhaseCommittedtist(l.ctx, &user.PhaseCommittedtistReq{
		Stuid:    req.Stuid,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PhaseCommittedtistResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
