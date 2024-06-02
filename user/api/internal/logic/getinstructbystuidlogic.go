package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetInstructByStuIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInstructByStuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInstructByStuIdLogic {
	return &GetInstructByStuIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInstructByStuIdLogic) GetInstructByStuId(req *types.GetInstructByStuIdReq) (*types.GetInstructByStuIdResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.GetInstructByStuId(l.ctx, &user.GetInstructByStuIdReq{
		StuId: req.StuID,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetInstructByStuIdResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
