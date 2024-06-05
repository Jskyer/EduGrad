package logic

import (
	"context"
	"fmt"

	"edu-grad/common"
	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhaseAllowedtistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhaseAllowedtistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhaseAllowedtistLogic {
	return &PhaseAllowedtistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhaseAllowedtistLogic) PhaseAllowedtist(in *user.PhaseAllowedtistReq) (*user.PhaseAllowedtistResp, error) {
	// todo: add your logic here and delete this line
	doneInstructs, err := l.svcCtx.InstructModel.ListByStuId(l.ctx, in.Stuid)
	if err != nil {
		return nil, err
	}

	fmt.Println("doneModel")
	fmt.Println(len(doneInstructs))

	var phaseRelMods []*model.PhaseRelation
	var totalPage int64 = 0
	var cnt int64 = 0

	if len(doneInstructs) > 0 {
		userids := make([]string, len(doneInstructs))
		for i, doneModel := range doneInstructs {
			userids[i] = doneModel.TeacherID
			fmt.Printf("%+v\n", doneModel)
		}

		phaseRelMods, totalPage, cnt, err = l.svcCtx.PhaseRelationModel.ListTeachersByPhaseidExceptTids(l.ctx, in.Phaseid, userids, in.PageNum, in.PageSize)
		if err != nil {
			return nil, err
		}

	} else {
		phaseRelMods, totalPage, cnt, err = l.svcCtx.PhaseRelationModel.ListByPhaseidAndIdentity(l.ctx, in.Phaseid, "导师", in.PageNum, in.PageSize)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("prmodel")
	fmt.Println(len(phaseRelMods))

	tarids := make([]string, len(phaseRelMods))
	for i, prmodel := range phaseRelMods {
		tarids[i] = prmodel.UserId
		fmt.Printf("%+v\n", prmodel)
	}

	userData, err := l.svcCtx.UserModel.ListByUserids(l.ctx, tarids)
	if err != nil {
		return nil, err
	}

	fmt.Println("userData")
	fmt.Println(len(userData))
	userInfos := make([]*user.UserInfo, len(userData))
	for i, udatum := range userData {
		userInfos[i] = common.ConvertUser(udatum)
		fmt.Printf("%+v\n", udatum)
	}

	return &user.PhaseAllowedtistResp{
		Users: userInfos,
		Total: totalPage,
		Count: cnt,
	}, nil
}
