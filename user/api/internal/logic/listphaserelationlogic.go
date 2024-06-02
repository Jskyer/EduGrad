package logic

import (
	"context"
	"errors"

	"edu-grad/user/api/internal/svc"
	"edu-grad/user/api/internal/types"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListPhaseRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPhaseRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPhaseRelationLogic {
	return &ListPhaseRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPhaseRelationLogic) ListPhaseRelation(req *types.ListPhaseRelationReq) (*types.ListPhaseRelationResp, error) {
	// todo: add your logic here and delete this line
	_, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("id不合法")
	}

	if req.Identity != "学生" && req.Identity != "导师" {
		return nil, errors.New("identity不合法")
	}

	if req.PageNum <= 0 || req.PageSize <= 0 {
		return nil, errors.New("pageNum或pageSize不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.ListPhaseRelation(l.ctx, &user.ListPhaseRelationReq{
		Id:       req.Id,
		Identity: req.Identity,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var resp types.ListPhaseRelationResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
