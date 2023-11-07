package line_repo

import "github.com/ouahochenchen/LLS/initialize"

var dbLine = initialize.MasterDb

type LineRepo interface {
	Create(tab LineResourceTab) (uint64, error)
	Update(tab LineResourceTab) (uint64, error)
	SelectById(lineId uint64) (*LineResourceTab, error)
}

type lineRepoImpl struct {
}

func NewLineRepo() LineRepo {
	return &lineRepoImpl{}
}

func (l *lineRepoImpl) Create(tab LineResourceTab) (uint64, error) {
	tx := dbLine.Create(tab)
	return tab.LineId, tx.Error
}
func (l *lineRepoImpl) Update(tab LineResourceTab) (uint64, error) {
	err := dbLine.Updates(tab).Error
	return tab.LineId, err
}
func (l *lineRepoImpl) SelectById(lineId uint64) (*LineResourceTab, error) {
	var result LineResourceTab
	err := dbLine.Find(&result).Where("line_id=?", lineId).Error
	return &result, err
}
