package reource

import (
	"github.com/ouahochenchen/LLS/internal/dal/repository/line_repo"
	"github.com/ouahochenchen/LLS/internal/dal/repository/site_repo"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
)

type ResourceDomain interface {
	IsExistResource(request *_go.ExistSiteLineRequest) (*_go.ExistSiteLineResponse, error)
}

type resourceDomainImpl struct {
	lineRepo line_repo.LineRepo
	siteRepo site_repo.SiteRepo
}

func NewResourceDomain(l line_repo.LineRepo, s site_repo.SiteRepo) ResourceDomain {
	return &resourceDomainImpl{
		lineRepo: l,
		siteRepo: s,
	}
}
func (e *resourceDomainImpl) IsExistResource(request *_go.ExistSiteLineRequest) (*_go.ExistSiteLineResponse, error) {
	var isExist = true
	if request.ResourceType == 1 {
		res1, err := e.siteRepo.SelectById(request.ResourceId)
		if err != nil {
			return nil, err
		}
		if res1.SiteId == 0 {
			isExist = false
		}
	}
	if request.ResourceType == 2 {
		res2, err := e.lineRepo.SelectById(request.ResourceId)
		if err != nil {
			return nil, err
		}
		if res2.LineId == 0 {
			isExist = false
		}
	}
	if request.NextType == 1 {
		res3, err := e.siteRepo.SelectById(request.NextId)
		if err != nil {
			return nil, err
		}
		if res3.SiteId == 0 {
			isExist = false
		}
	}
	if request.NextType == 2 {
		res4, err := e.lineRepo.SelectById(request.NextId)
		if err != nil {
			return nil, err
		}
		if res4.LineId == 0 {
			isExist = false
		}
	}
	if (request.NextType != 1 && request.NextType != 2) || (request.ResourceType != 1 && request.ResourceType != 2) {
		isExist = false
	}
	return &_go.ExistSiteLineResponse{
		IsExist: isExist,
	}, nil
}
