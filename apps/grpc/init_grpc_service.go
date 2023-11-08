package grpc

import (
	"github.com/ouahochenchen/LLS/internal/dal/repository/line_repo"
	"github.com/ouahochenchen/LLS/internal/dal/repository/order_repo"
	"github.com/ouahochenchen/LLS/internal/dal/repository/site_repo"
	"github.com/ouahochenchen/LLS/internal/domain/order"
	"github.com/ouahochenchen/LLS/internal/domain/reource"
	"google.golang.org/grpc"
)

var GrpcServer *grpc.Server
var LineRepo line_repo.LineRepo
var SiteRepo site_repo.SiteRepo
var OrderRepo order_repo.OrderRepo
var ResourceDomain reource.ResourceDomain
var OrderDomain order.OrderDomain

func init() {
	GrpcServer = grpc.NewServer()
	LineRepo = line_repo.NewLineRepo()
	SiteRepo = site_repo.NewSiteRepo()
	OrderRepo = order_repo.NewOrderRepo()
	ResourceDomain = reource.NewResourceDomain(LineRepo, SiteRepo)
	OrderDomain =order.NewOderDomain(OrderRepo)
}
