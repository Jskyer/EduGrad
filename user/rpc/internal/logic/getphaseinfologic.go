package logic

import (
	"context"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPhaseinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPhaseinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPhaseinfoLogic {
	return &GetPhaseinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPhaseinfoLogic) GetPhaseinfo(in *user.PhaseInfoReq) (*user.PhaseInfoResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.PhaseModel.FindOne(l.ctx, in.Phaseid)

	if err != nil {
		return nil, err
	}

	return &user.PhaseInfoResp{
		Id:      data.ID.Hex(),
		Term:    data.Term,
		Process: data.Process,
	}, nil
}
