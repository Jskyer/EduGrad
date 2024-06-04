package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTermSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTermSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTermSortLogic {
	return &ListTermSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTermSortLogic) ListTermSort(in *user.ListTermSortReq) (*user.ListTermSortResp, error) {
	// todo: add your logic here and delete this line
	phaseModels, totalPage, cnt, err := l.svcCtx.PhaseModel.ListByTermSorted(l.ctx, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	phases := make([]*user.Phase, len(phaseModels))
	for i, phaseModel := range phaseModels {
		phases[i] = common.ConvertPhase(phaseModel)
	}

	return &user.ListTermSortResp{
		Phases: phases,
		Total:  totalPage,
		Count:  cnt,
	}, nil
}
