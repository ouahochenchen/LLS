package grpc

func InitRouter() {
	ResourceGrpcHandler("tcp", "50051")
}
