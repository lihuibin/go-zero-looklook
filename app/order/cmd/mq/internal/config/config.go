package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	//kq
	PaymentUpdateStatusConf kq.KqConf

	//rpc
	OrderRpcConf      zrpc.RpcClientConf
	MqueueRpcConf     zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
}
