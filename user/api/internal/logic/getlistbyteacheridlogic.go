package logic

import (
	"context"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetListByTeacherIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListByTeacherIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByTeacherIdLogic {
	return &GetListByTeacherIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListByTeacherIdLogic) GetListByTeacherId(req *types.GetListByTeacherIdReq) (*types.GetListByTeacherIdResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.GetListByTeacherId(l.ctx, &user.GetListByTeacherIdReq{
		TeacherId: req.TeacherID,
	})
	if err != nil {
		return nil, err
	}

	var resp types.GetListByTeacherIdResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
