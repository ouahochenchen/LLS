package order

import (
	"github.com/ouahochenchen/LLS/internal/dal/repository/order_repo"
	"github.com/ouahochenchen/LLS/internal/util"
	_go "github.com/ouahochenchen/LLS/protocol/grpc/go"
	"strconv"
)

type OrderDomain interface {
	PlacingOrder(request *_go.LfsRequest) (*_go.LfsResponse, error)
}
type orderDomainImpl struct {
	orderRepo order_repo.OrderRepo
}

func NewOderDomain(orderRepo order_repo.OrderRepo) OrderDomain {
	return &orderDomainImpl{
		orderRepo: orderRepo,
	}
}
func (o *orderDomainImpl) PlacingOrder(request *_go.LfsRequest) (*_go.LfsResponse, error) {
	result, err := o.orderRepo.SelectByLaneOrderId(request.LfsOrderId)
	if err != nil {
		return nil, err
	}
	if result.OrderId != 0 {
		return nil, &util.MyError{MyErrorMessage: "已经有重复订单"}
	}

	var isDeliverFunc = func() bool {
		return true
	}

	if !isDeliverFunc() {
		return nil, &util.MyError{MyErrorMessage: "链路不可达"}
	}
	worker := util.NewIdWorker(int64(request.LfsOrderId))
	orderId := worker.NextId()
	createOrderId, err := o.orderRepo.Create(
		&order_repo.LineOrderTab{
			OrderId:       uint64(orderId),
			BuyerName:     request.BuyerName,
			BuyerAddress:  request.BuyerAddress,
			BuyerPhone:    request.BuyerPhone,
			GoodsType:     request.GoodsType,
			SellerName:    request.SellerName,
			SellerAddress: request.SellerAddress,
			SellerPhone:   request.SellerPhone,
			PackageHeight: request.PackageHeight,
			PackageWeight: request.PackageWeight,
			Price:         float64(request.Price),
			OrderStatus:   request.OrderStatus,
			LineId:        request.LineId,
			LaneOrderId:   request.LfsOrderId,
		},
	)
	if err != nil {
		return nil, err
	}
	return &_go.LfsResponse{
		LfsOrderId:  strconv.FormatUint(request.LfsOrderId, 10),
		OrderId:     strconv.FormatUint(createOrderId, 10),
		OrderStatus: strconv.FormatUint(request.OrderStatus, 10),
		IsOk:        true,
	}, nil
}
