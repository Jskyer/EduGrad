package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListByTeacherIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListByTeacherIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByTeacherIdLogic {
	return &GetListByTeacherIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListByTeacherIdLogic) GetListByTeacherId(in *user.GetListByTeacherIdReq) (*user.GetListByTeacherIdResp, error) {
	// todo: add your logic here and delete this line
	instructModels, err := l.svcCtx.InstructModel.ListByTeacherId(l.ctx, in.TeacherId)
	if err != nil {
		return nil, err
	}

	instructs := make([]*user.InstructRelation, len(instructModels))
	for i, instructModel := range instructModels {
		instructs[i] = common.ConvertInstruct(instructModel)
	}

	// for _, instructModel := range instructModels {
	// 	fmt.Printf("model: %+v\n", instructModel)
	// }

	// fmt.Println("============")
	// for _, resModel := range instructs {
	// 	fmt.Printf("res: %+v\n", resModel)
	// }

	return &user.GetListByTeacherIdResp{
		Instructs: instructs,
	}, nil
}
