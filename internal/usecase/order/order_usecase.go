package order

import (
	"context"
	"github.com/ouahochenchen/LLS/internal/domain/order"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
	"google.golang.org/grpc"
)

type OrderUseCase struct {
	_go.UnimplementedLfsServiceServer
	domain order.OrderDomain
}

func NewOrderUseCase(domain order.OrderDomain) *OrderUseCase {
	return &OrderUseCase{
		domain: domain,
	}
}
func (o *OrderUseCase) CreateLineOrder(ctx context.Context, req *_go.LfsRequest, opts ...grpc.CallOption) (*_go.LfsResponse, error) {
	resp, err := o.domain.PlacingOrder(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
