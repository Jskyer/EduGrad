package svc

import (
	"edu-grad/user/model"
	"edu-grad/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	RdsCli    *redis.Redis

	InstructModel      model.InstructRelationModel
	PhaseModel         model.PhaseModel
	PhaseRelationModel model.PhaseRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		UserModel:          model.NewUserModel(c.Mongo.URL, c.Mongo.DB, c.Mongo.CollectionName),
		RdsCli:             redis.MustNewRedis(c.Redis.RedisConf),
		InstructModel:      model.NewInstructRelationModel(c.Mongo.URL, c.Mongo.DB, "instruct_relation"),
		PhaseModel:         model.NewPhaseModel(c.Mongo.URL, c.Mongo.DB, "phase"),
		PhaseRelationModel: model.NewPhaseRelationModel(c.Mongo.URL, c.Mongo.DB, "phase_relation"),
	}
}
