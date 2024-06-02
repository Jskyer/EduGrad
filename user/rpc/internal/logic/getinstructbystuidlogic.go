package logic

import (
	"context"

	"edu-grad/common"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInstructByStuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInstructByStuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInstructByStuIdLogic {
	return &GetInstructByStuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInstructByStuIdLogic) GetInstructByStuId(in *user.GetInstructByStuIdReq) (*user.GetInstructByStuIdResp, error) {
	// todo: add your logic here and delete this line
	instructModels, err := l.svcCtx.InstructModel.ListByStuId(l.ctx, in.StuId)
	if err != nil {
		return nil, err
	}

	instructs := make([]*user.InstructRelation, len(instructModels))
	for i, instructModel := range instructModels {
		instructs[i] = common.ConvertInstruct(instructModel)
	}

	// var res []*user.InstructRelation
	// _ = copier.Copy(&res, instructModels)

	// fmt.Printf("instructModel: %+v\n", instructModel)
	// fmt.Printf("res: %+v\n", res)
	// for _, instructModel := range instructModels {
	// 	fmt.Printf("model: %+v\n", instructModel)
	// }

	// fmt.Println("============")
	// for _, resModel := range instructs {
	// 	fmt.Printf("res: %+v\n", resModel)
	// }

	return &user.GetInstructByStuIdResp{
		Instructs: instructs,
	}, nil
}
