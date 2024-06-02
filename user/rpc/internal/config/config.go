package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mongo struct {
		URL            string
		DB             string
		CollectionName string
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
