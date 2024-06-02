package logic

import (
	"context"
	"errors"
	"strconv"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStateLogic {
	return &UpdateStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStateLogic) UpdateState(in *user.UpdateStateReq) (*user.UpdateStateResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.InstructModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	validState, err := strconv.Atoi(in.State)
	if err != nil {
		return nil, errors.New("state不合法")
	}

	if validState < 1 || validState > 7 {
		return nil, errors.New("state不合法")
	}

	updateResult, err := l.svcCtx.InstructModel.Update(l.ctx, &model.InstructRelation{
		ID:    oid,
		State: in.State,
	})
	if err != nil {
		return nil, err
	}

	return &user.UpdateStateResp{
		Updated: updateResult.ModifiedCount,
	}, nil
}
