package svc

import (
	"edu-grad/paper/api/internal/config"
	"edu-grad/paper/rpc/paperclient"
	"edu-grad/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PaperRpc paperclient.Paper
	UserRpc  userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		PaperRpc: paperclient.NewPaper(zrpc.MustNewClient(c.PaperRpc)),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
