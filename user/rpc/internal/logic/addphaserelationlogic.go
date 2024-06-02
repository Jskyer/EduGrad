package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type AddPhaseRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhaseRelationLogic {
	return &AddPhaseRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPhaseRelationLogic) AddPhaseRelation(in *user.AddPhaseRelationReq) (*user.AddPhaseRelationResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.PhaseModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 判断phaserelation是否存在该phaseid
	_, err = l.svcCtx.PhaseRelationModel.FindOneByPhaseid(l.ctx, in.Id)
	if err != mon.ErrNotFound {
		fmt.Println("这期phase已经添加过用户")
		return nil, errors.New("这期phase已经添加过用户")
	}

	useridSet := mapset.NewSet[string]()

	if in.Grade != "" {
		fmt.Println("Grade != 空")
		usersByGrade, err := l.svcCtx.UserModel.ListTeacherOrStuByGrade(l.ctx, in.Grade)

		if err != nil {
			return nil, err
		}

		for _, userByGrade := range usersByGrade {
			useridSet.Add(userByGrade.ID.Hex())
		}
	}

	if len(in.Users) > 0 {
		fmt.Println("Users != 空")
		for _, userInReq := range in.Users {
			// oid, err := primitive.ObjectIDFromHex(userInReq)

			// if err != nil {
			// 	return nil, errors.New("users不合法")
			// }

			useridSet.Add(userInReq)
		}
	}

	fmt.Printf("Cardinality: %d\n", useridSet.Cardinality())
	userids := make([]string, useridSet.Cardinality())
	idx := 0
	for userid := range useridSet.Iter() {
		userids[idx] = userid
		idx++
	}

	userIdsSamePhaseid, err := l.svcCtx.PhaseRelationModel.ExistUseridsInPhaseRelation(l.ctx, in.Id, userids)
	fmt.Println(userIdsSamePhaseid)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return nil, err
	}
	if len(userIdsSamePhaseid) > 0 {
		fmt.Println("存在phaserelation,userid和phaseid相同")
		return nil, errors.New("存在phaserelation,userid和phaseid相同")
	}

	usermodels, err := l.svcCtx.UserModel.ListByUserids(l.ctx, userids)
	if err != nil {
		fmt.Println("ListByUserids err")
		return nil, err
	}

	phaseRelations := make([]*model.PhaseRelation, len(usermodels))
	for i, usermodel := range usermodels {
		phaseRelations[i] = &model.PhaseRelation{
			UserId:   usermodel.ID.Hex(),
			Identity: usermodel.Identity,
			PhaseId:  in.Id,
		}
	}

	err = l.svcCtx.PhaseRelationModel.InsertMany(l.ctx, phaseRelations)
	if err != nil {
		fmt.Println("InsertMany err")
		return nil, err
	}

	return &user.AddPhaseRelationResp{
		Ok: 1,
	}, nil
}
