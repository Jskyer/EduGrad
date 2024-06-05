package logic

import (
	"context"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhaseCommittedtistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhaseCommittedtistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhaseCommittedtistLogic {
	return &PhaseCommittedtistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhaseCommittedtistLogic) PhaseCommittedtist(in *user.PhaseCommittedtistReq) (*user.PhaseCommittedtistResp, error) {
	// todo: add your logic here and delete this line
	instructModels, totalPage, cnt, err := l.svcCtx.InstructModel.ListPageByStuId(l.ctx, in.Stuid, in.PageNum, in.PageSize)

	if err != nil {
		return nil, err
	}

	userids := make([]string, len(instructModels))
	for i, instructModel := range instructModels {
		userids[i] = instructModel.TeacherID
	}

	userData, err := l.svcCtx.UserModel.ListByUserids(l.ctx, userids)
	if err != nil {
		return nil, err
	}

	instructUsers := make([]*user.InstructUser, len(userData))
	for i, umodel := range userData {
		curPermit := ""

		for _, imodel := range instructModels {
			if imodel.TeacherID == umodel.ID.Hex() {
				curPermit = imodel.Permit
				break
			}
		}

		instructUsers[i] = &user.InstructUser{
			Id:       umodel.ID.Hex(),
			Username: umodel.Username,
			Identity: umodel.Identity,
			Grade:    umodel.Grade,
			Permit:   curPermit,
		}
	}

	return &user.PhaseCommittedtistResp{
		Users: instructUsers,
		Total: totalPage,
		Count: cnt,
	}, nil
}
