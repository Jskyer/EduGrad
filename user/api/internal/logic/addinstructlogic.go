package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddInstructLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddInstructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInstructLogic {
	return &AddInstructLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInstructLogic) AddInstruct(req *types.AddInstructReq) (*types.AddInstructResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.AddInstruct(l.ctx, &user.AddInstructReq{
		StuId:     req.StuID,
		TeacherId: req.TeacherID,
		Permit:    req.Permit,
	})

	if err != nil {
		return nil, err
	}

	var resp types.AddInstructResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
