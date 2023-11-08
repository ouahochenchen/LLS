package grpc

import (
	"github.com/ouahochenchen/LLS/internal/dal/repository/line_repo"
	"github.com/ouahochenchen/LLS/internal/dal/repository/site_repo"
	"github.com/ouahochenchen/LLS/internal/domain/reource"
	"google.golang.org/grpc"
)

var GrpcServer *grpc.Server
var LineRepo line_repo.LineRepo
var SiteRepo site_repo.SiteRepo
var ResourceDomain reource.ResourceDomain

func init() {
	GrpcServer = grpc.NewServer()
	LineRepo = line_repo.NewLineRepo()
	SiteRepo = site_repo.NewSiteRepo()
	ResourceDomain = reource.NewResourceDomain(LineRepo, SiteRepo)
}
