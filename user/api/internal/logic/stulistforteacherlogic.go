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

type StulistForTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStulistForTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StulistForTeacherLogic {
	return &StulistForTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StulistForTeacherLogic) StulistForTeacher(req *types.StulistForTeacherReq) (*types.StulistForTeacherResp, error) {
	// todo: add your logic here and delete this line
	if req.PageNum <= 0 || req.PageSize <= 0 {
		return nil, errors.New("pageNum或pageSize不合法")
	}

	if req.Permit != "0" && req.Permit != "1" && req.Permit != "2" {
		return nil, errors.New("permit不合法")
	}

	_, err := primitive.ObjectIDFromHex(req.Teacherid)
	if err != nil {
		return nil, errors.New("teacherid不合法")
	}

	rpcResp, err := l.svcCtx.UserRpc.StulistForTeacher(l.ctx, &user.StulistForTeacherReq{
		Teacherid: req.Teacherid,
		Permit:    req.Permit,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var resp types.StulistForTeacherResp
	err = copier.Copy(&resp, rpcResp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
