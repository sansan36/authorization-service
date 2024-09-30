package main

import (
	"github.com/sansan36/authorization-service/core"
	userv1 "github.com/sansan36/authorization-service/gen/authorization/v1"
	"github.com/sansan36/authorization-service/gen/authorization/v1/authorizationv1connect"
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
	path, handler := authorizationv1connect.NewAuthorizationServiceHandler(
		serviceHandler,
		core.NewInterceotors(),
	)

	// run the server
	core.RunServer(core.ServerSpec{
		RpcPath:                            path,
		RpcHandler:                         handler,
		RegisterServiceHandlerFromEndpoint: userv1.RegisterAuthorizationServiceHandlerFromEndpoint,
		RpcPort:                            rpcPort,
		HttpPort:                           httpPort,
	})

}
