package order_repo

import "github.com/ouahochenchen/LLS/initialize"

var dbOrder = initialize.MasterDb

type OrderRepo interface {
	Create(tab *LineOrderTab) (uint64, error)
	Update(tab *LineOrderTab) (uint64, error)
	SelectByLaneOrderId(id uint64) (*LineOrderTab, error)
	SelectByLineId(id uint64) (*LineOrderTab, error)
	SelectByOrderId(id uint64) (*LineOrderTab, error)
}
type orderRepoImpl struct {
}

func NewOrderRepo() OrderRepo {
	return &orderRepoImpl{}
}

func (o *orderRepoImpl) Create(tab *LineOrderTab) (uint64, error) {
	err := dbOrder.Create(tab).Error
	return tab.OrderId, err
}
func (o *orderRepoImpl) Update(tab *LineOrderTab) (uint64, error) {
	err := dbOrder.Updates(tab).Error
	return tab.OrderId, err
}
func (o *orderRepoImpl) SelectByLaneOrderId(id uint64) (*LineOrderTab, error) {
	var result LineOrderTab
	err := dbOrder.Find(&result).Where("lane_order_id=?", id).Error
	return &result, err
}
func (o *orderRepoImpl) SelectByLineId(id uint64) (*LineOrderTab, error) {
	var result LineOrderTab
	err := dbOrder.Find(&result).Where("line_id=?", id).Error
	return &result, err
}
func (o *orderRepoImpl) SelectByOrderId(id uint64) (*LineOrderTab, error) {
	var result LineOrderTab
	err := dbOrder.Find(&result).Where("order_id=?", id).Error
	return &result, err
}
