package main

import (
	"github.com/ouahochenchen/LLS/apps/grpc"
	_ "github.com/ouahochenchen/LLS/initialize"
)

func main() {
	grpc.InitRouter()
}
