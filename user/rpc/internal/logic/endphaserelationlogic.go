package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type EndPhaseRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEndPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndPhaseRelationLogic {
	return &EndPhaseRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EndPhaseRelationLogic) EndPhaseRelation(in *user.EndPhaseRelationReq) (*user.EndPhaseRelationResp, error) {
	// 查询所有该phaseid中是学生的relation
	phaseRelationMods, err := l.svcCtx.PhaseRelationModel.ListAllByPhaseidAndIdentity(l.ctx, in.Id, "学生")
	if err != nil {
		return nil, err
	}

	fmt.Println("=========")
	stuIds := make([]string, len(phaseRelationMods))
	for i, pmod := range phaseRelationMods {
		stuIds[i] = pmod.UserId
		fmt.Printf("%+v\n", pmod)
	}
	fmt.Println("=========")

	// 从InstructRelation 查询终稿通过的学生的记录
	instructModels, err := l.svcCtx.InstructModel.ListPassedByStuids(l.ctx, stuIds)
	if err != nil {
		return nil, err
	}

	passedStuidSet := mapset.NewSet[string]()
	for _, instructModel := range instructModels {
		passedStuidSet.Add(instructModel.StuID)
	}

	var passedPhaseRelations []string
	var failedPhaseRelations []string

	// 遍历phaseRelation，分出通过、不通过的两批slice
	for _, pmod := range phaseRelationMods {
		if passedStuidSet.ContainsOne(pmod.UserId) {
			// pmod.Pass = 1
			passedPhaseRelations = append(passedPhaseRelations, pmod.UserId)
		} else {
			// pmod.Pass = 2
			failedPhaseRelations = append(failedPhaseRelations, pmod.UserId)
		}
	}

	fmt.Println("=========pass")
	fmt.Printf("%d\n", len(passedPhaseRelations))
	for _, pmod := range passedPhaseRelations {
		fmt.Println(pmod)
	}

	fmt.Println("=========fail")
	fmt.Printf("%d\n", len(failedPhaseRelations))
	for _, pmod := range failedPhaseRelations {
		fmt.Println(pmod)
	}

	fmt.Println("=========")

	if len(passedPhaseRelations) > 0 {
		passedUpt, err := l.svcCtx.PhaseRelationModel.UpdateStudentsPass(l.ctx, in.Id, passedPhaseRelations, 1)
		if err != nil {
			return nil, errors.New("update pass error")
		}
		if int(passedUpt.ModifiedCount) != len(passedPhaseRelations) {
			return nil, errors.New("update pass result error")
		}
	}

	if len(failedPhaseRelations) > 0 {
		failedUpt, err := l.svcCtx.PhaseRelationModel.UpdateStudentsPass(l.ctx, in.Id, failedPhaseRelations, 2)
		if err != nil {
			return nil, errors.New("update fail error")
		}
		if int(failedUpt.ModifiedCount) != len(failedPhaseRelations) {
			return nil, errors.New("update fail result error")
		}
	}

	return &user.EndPhaseRelationResp{
		Ok: 1,
	}, nil
}
