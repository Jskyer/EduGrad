package logic

import (
	"context"
	"errors"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCondPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserCondPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCondPageLogic {
	return &GetUserCondPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCondPageLogic) GetUserCondPage(req *types.GetUserCondPageReq) (*types.GetUserCondPageResp, error) {
	// todo: add your logic here and delete this line
	if req.PageNum <= 0 || req.PageSize <= 0 {
		return nil, errors.New("pageNum或pageSize不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.GetUserCondPage(l.ctx, &user.GetUserCondPageReq{
		Grade:    req.Grade,
		Identity: req.Identity,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetUserCondPageResp
	_ = copier.Copy(&resp, rpcResp)

	return &resp, nil
}
