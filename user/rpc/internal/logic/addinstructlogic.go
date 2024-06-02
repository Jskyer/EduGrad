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

type AddInstructLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInstructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInstructLogic {
	return &AddInstructLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddInstructLogic) AddInstruct(in *user.AddInstructReq) (*user.AddInstructResp, error) {
	// todo: add your logic here and delete this line
	if in.Permit != "0" {
		logx.Info("permit 不合法")
		return nil, errors.New("permit 不合法")
	}

	// 确认stuId 当前不存在permit=1 instruct
	permited, err := l.svcCtx.InstructModel.FindPermitedByStuId(l.ctx, in.StuId)
	if permited != nil {
		logx.Info("err 1 exist permit")
		return nil, errors.New("已存在通过的指导关系，不可再申请")
	}

	if err != mon.ErrNotFound {
		logx.Info("err 2")
		return nil, err
	}

	// 确认instruct 不存在同stuId, teacherId
	exist, err := l.svcCtx.InstructModel.FindOneByStuIdTeacherId(l.ctx, in.StuId, in.TeacherId)
	if exist != nil {
		logx.Info("err 3")
		return nil, errors.New("向同一个导师申请过指导关系，不可再申请")
	}

	if err != mon.ErrNotFound {
		logx.Info("err 4")
		return nil, err
	}

	stuModel, err := l.svcCtx.UserModel.FindStuById(l.ctx, in.StuId)
	if err != nil {
		logx.Info("err 5")
		return nil, err
	}

	teacherModel, err := l.svcCtx.UserModel.FindTeacherById(l.ctx, in.TeacherId)
	if err != nil {
		logx.Info("err 6")
		return nil, err
	}

	err = l.svcCtx.InstructModel.Insert(l.ctx, &model.InstructRelation{
		StuID:       in.StuId,
		StuName:     stuModel.Username,
		TeacherID:   in.TeacherId,
		TeacherName: teacherModel.Username,
		Permit:      in.Permit,
	})

	if err != nil {
		logx.Info("err 7")
		return nil, err
	}

	return &user.AddInstructResp{
		Ok: 1,
	}, nil
}
