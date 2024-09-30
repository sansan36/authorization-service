package main

import (
	"github.com/sansan36/authorization-service/core"
	userv1 "github.com/sansan36/authorization-service/gen/user/v1"
	"github.com/sansan36/authorization-service/gen/user/v1/userv1connect"
	"github.com/sansan36/authorization-service/handler"
	"github.com/sansan36/authorization-service/repository"
)

func main() {
	var rpcPort int = 9091
	var httpPort int = 8081

	// start DB connection
	core.StartDBConnection()
	defer core.CloseDBMain()

	// initiate RPC path and handler
	serviceHandler := handler.NewUserHandler(
		repository.NewUserRepository(core.DBMain),
	)
	path, handler := userv1connect.NewUserServiceHandler(
		serviceHandler,
		core.NewInterceotors(),
	)

	// run the server
	core.RunServer(core.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: userv1.RegisterUserServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

}
