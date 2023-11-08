package resource

import (
	"context"
	"github.com/ouahochenchen/LLS/internal/domain/reource"

	//"github.com/ouahochenchen/LLS/apps/grpc"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

//type ResourceUsecase interface {
//	IsExistResource(ctx context.Context, req *_go.ExistSiteLineRequest) (*_go.ExistSiteLineResponse, error)
//}

type ResourceUseCase struct {
	_go.UnimplementedSiteServiceServer
	resourceDomain reource.ResourceDomain
}

func NewResourceUsecase(resourceDomain reource.ResourceDomain) *ResourceUseCase {
	return &ResourceUseCase{
		resourceDomain: resourceDomain,
	}
}

func (r *ResourceUseCase) IsExistResource(ctx context.Context, req *_go.ExistSiteLineRequest) (*_go.ExistSiteLineResponse, error) {
	//req := info.(*_go.ExistSiteLineRequest)
	resp, err := r.resourceDomain.IsExistResource(req)
	return resp, err
}
