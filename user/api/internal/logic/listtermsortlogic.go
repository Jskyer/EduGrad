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

type ListTermSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTermSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTermSortLogic {
	return &ListTermSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTermSortLogic) ListTermSort(req *types.ListTermSortReq) (*types.ListTermSortResp, error) {
	// todo: add your logic here and delete this line
	if req.PageNum <= 0 || req.PageSize <= 0 {
		return nil, errors.New("pageNum或pageSize不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.ListTermSort(l.ctx, &user.ListTermSortReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var resp types.ListTermSortResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
