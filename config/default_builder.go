package config

import (
	"context"
	"fmt"
	"gim/pkg/grpclib/picker"
	_ "gim/pkg/grpclib/resolver/addrs"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type defaultBuilder struct{}

func (*defaultBuilder) Build() Configuration {
	logger.Level = zap.DebugLevel
	logger.Target = logger.Console

	return Configuration{
		MySQL:                "root:Gols852456mysql@tcp(rm-bp1pq4p15633m4n9ceo.mysql.rds.aliyuncs.com:3306)/im?charset=utf8mb4&parseTime=True&loc=Local",
		RedisHost:            "82.156.174.11:6379",
		RedisPassword:        "gols852456redis.",
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		ConnectLocalAddr:     "localhost:8000",
		ConnectRPCListenAddr: ":8000",
		ConnectTCPListenAddr: ":8001",
		ConnectWSListenAddr:  ":8002",

		LogicRPCListenAddr:    ":8010",
		BusinessRPCListenAddr: ":8020",
		FileHTTPListenAddr:    "8030",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///localhost:8000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///localhost:8010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		BusinessIntClientBuilder: func() pb.BusinessIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///localhost:8010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessIntClient(conn)
		},
	}
}
