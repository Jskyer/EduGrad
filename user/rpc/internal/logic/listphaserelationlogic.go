package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPhaseRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPhaseRelationLogic {
	return &ListPhaseRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPhaseRelationLogic) ListPhaseRelation(in *user.ListPhaseRelationReq) (*user.ListPhaseRelationResp, error) {
	// todo: add your logic here and delete this line
	phaseRelationModels, totalPage, err := l.svcCtx.PhaseRelationModel.ListByPhaseidAndIdentity(l.ctx, in.Id, in.Identity, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	userids := make([]string, len(phaseRelationModels))
	for i, pmodel := range phaseRelationModels {
		userids[i] = pmodel.UserId
	}

	users, err := l.svcCtx.UserModel.ListByUserids(l.ctx, userids)
	if err != nil {
		return nil, err
	}

	userinfos := make([]*user.UserInfo, len(users))
	for i, umodel := range users {
		userinfos[i] = common.ConvertUser(umodel)
	}

	return &user.ListPhaseRelationResp{
		Users: userinfos,
		Total: totalPage,
	}, nil
}
