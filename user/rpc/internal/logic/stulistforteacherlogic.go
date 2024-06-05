package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type StulistForTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStulistForTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StulistForTeacherLogic {
	return &StulistForTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StulistForTeacherLogic) StulistForTeacher(in *user.StulistForTeacherReq) (*user.StulistForTeacherResp, error) {
	// todo: add your logic here and delete this line
	instructModels, totalPage, cnt, err := l.svcCtx.InstructModel.ListByTeacherIdPermit(l.ctx, in.Teacherid, in.Permit, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	stuids := make([]string, len(instructModels))
	for i, instructModel := range instructModels {
		stuids[i] = instructModel.StuID
	}

	userData, err := l.svcCtx.UserModel.ListByUserids(l.ctx, stuids)
	if err != nil {
		return nil, err
	}

	userDto := make([]*user.UserInfo, len(userData))
	for i, umodel := range userData {
		userDto[i] = common.ConvertUser(umodel)
	}

	return &user.StulistForTeacherResp{
		Users: userDto,
		Total: totalPage,
		Count: cnt,
	}, nil
}
