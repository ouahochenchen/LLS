package order

import (
	"context"
	"github.com/ouahochenchen/LLS/internal/domain/order"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
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
func (o *OrderUseCase) LfsRpc(ctx context.Context, req *_go.LfsRequest) (*_go.LfsResponse, error) {
	resp, err := o.domain.PlacingOrder(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
