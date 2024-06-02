package logic

import (
	"context"
	"errors"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type AddPhaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPhaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhaseLogic {
	return &AddPhaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPhaseLogic) AddPhase(in *user.AddPhaseReq) (*user.AddPhaseResp, error) {
	// todo: add your logic here and delete this line
	exist, err := l.svcCtx.PhaseModel.FindOneByTerm(l.ctx, in.Term)
	if exist != nil {
		return nil, errors.New("term已存在")
	}

	if err != mon.ErrNotFound {
		return nil, err
	}

	if in.Process != 1 {
		return nil, errors.New("process不合法")
	}

	err = l.svcCtx.PhaseModel.Insert(l.ctx, &model.Phase{
		Term:    in.Term,
		Process: in.Process,
	})
	if err != nil {
		return nil, err
	}

	return &user.AddPhaseResp{
		Ok: 1,
	}, nil
}
