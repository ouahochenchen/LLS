package grpc

func InitOrderRouter() {
	OrderGrpcHandler("tcp", "50052")
}
func InitResourceRouter() {
	ResourceGrpcHandler("tcp", "50051")
}
