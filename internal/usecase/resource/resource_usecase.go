package resource

import (
	"context"
	"github.com/ouahochenchen/LLS/apps/api"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

type ResourceUseCase struct {
	_go.UnimplementedSiteServiceServer
}

func (*ResourceUseCase) IsExistResource(ctx context.Context, req *_go.ExistSiteLineRequest) (*_go.ExistSiteLineResponse, error) {
	//req := info.(*_go.ExistSiteLineRequest)
	resp, err := api.ResourceDomain.IsExistResource(req)
	return resp, err
}
