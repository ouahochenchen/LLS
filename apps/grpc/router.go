package grpc

func InitRouter() {
	ResourceGrpcHandler("tcp", "50051")
	OrderGrpcHandler("tcp", "50052")
}
