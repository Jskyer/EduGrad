package logic

import (
	"context"
	"errors"
	"fmt"

	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/svc"
	"edu-grad/user/rpc/types/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateProcessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProcessLogic {
	return &UpdateProcessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProcessLogic) UpdateProcess(in *user.UpdateProcessReq) (*user.UpdateProcessResp, error) {
	// todo: add your logic here and delete this line
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	currentPhase, err := l.svcCtx.PhaseModel.FindCurrentPhase(l.ctx)
	if currentPhase != nil {
		currentPhase.Process = 1
		upt, err := l.svcCtx.PhaseModel.Update(l.ctx, currentPhase)

		if err != nil || upt.ModifiedCount != 1 {
			return nil, errors.New("original phase update error")
		}

		fmt.Printf("upt: %d\n", upt.ModifiedCount)
	}

	if err != nil && err != mon.ErrNotFound {
		return nil, err
	}

	updateResult, err := l.svcCtx.PhaseModel.Update(l.ctx, &model.Phase{
		ID:      oid,
		Process: in.Process,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("updateResult: %d\n", updateResult.ModifiedCount)

	return &user.UpdateProcessResp{
		Updated: updateResult.ModifiedCount,
	}, nil
}
