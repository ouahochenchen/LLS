package grpc

import (
	"github.com/ouahochenchen/LLS/internal/usecase/resource"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
	"log"
	"net"
)

//type SimpleHandlerFunc func(ctx context.Context, info interface{}) (interface{}, error)
//type SetPort func(protocol string, port string, ctx context.Context)

func ResourceGrpcHandler(protocol string, port string) {
	//request := info.(*_go.ExistSiteLineRequest)

	lis, err := net.Listen(protocol, ":"+port)
	if err != nil {
		log.Fatalf("无法监听：%v", err)
	}
	//s := grpc.NewServer()
	//pb.RegisterYourServiceServer(s, &server{})

	_go.RegisterSiteServiceServer(GrpcServer, resource.NewResourceUsecase(ResourceDomain))
	log.Println("启动 gRPC 服务...")

	if err := GrpcServer.Serve(lis); err != nil {
		log.Fatalf("无法启动服务：%v", err)
	}

}
