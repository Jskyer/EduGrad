package logic

import (
	"context"
	"fmt"

	"edu-grad/common"
	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (*types.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	userId := common.GetUidFromCtx(l.ctx)
	fmt.Println("userId:", userId)
	// if userId == "" {
	// 	fmt.Println("GetUidFromCtx error")
	// 	return nil, errors.New("GetUidFromCtx error")
	// }

	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetUserInfoResp
	_ = copier.Copy(&resp, rpcResp)
	fmt.Printf("resp: %+v\n", resp)

	return &resp, nil
}
