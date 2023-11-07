package site_repo

import "github.com/ouahochenchen/LLS/initialize"

var dbls = initialize.MasterDb

type SiteRepo interface {
	Create(tab SiteResourceTab) (uint64, error)
	Update(tab SiteResourceTab) (uint64, error)
	SelectById(siteId uint64) (*SiteResourceTab, error)
}

type siteRepoImpl struct {
}

func (s *siteRepoImpl) Create(tab SiteResourceTab) (uint64, error) {
	db := dbls.Create(tab)
	if db.Error != nil {
		return 0, db.Error
	}
	return tab.SiteId, nil
}
func NewSiteRepo() SiteRepo {
	return &siteRepoImpl{}
}
func (s *siteRepoImpl) Update(tab SiteResourceTab) (uint64, error) {
	db := dbls.Updates(tab)
	if db.Error != nil {
		return 0, db.Error
	}
	return tab.SiteId, nil
}
func (s *siteRepoImpl) SelectById(siteId uint64) (*SiteResourceTab, error) {
	result := SiteResourceTab{}
	err := dbls.Find(&result).Where("site_id=?", siteId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
