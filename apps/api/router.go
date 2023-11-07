package api

import "github.com/ouahochenchen/LLS/apps/handler"

func InitRouter() {
	handler.GrpcHandler("tcp", "50051")
}
