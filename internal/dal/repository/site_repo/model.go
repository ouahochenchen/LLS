package site_repo

type SiteResourceTab struct {
	SiteId           uint64 `gorm:"column:site_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"site_id"`
	SiteName         string `gorm:"column:site_name;type:varchar(255)" json:"site_name"`
	RegionCode       string `gorm:"column:region_code;type:varchar(255)" json:"region_code"`
	SiteType         uint64 `gorm:"column:site_type;type:tinyint(4);default:0" json:"site_type"`
	SiteSubType      uint64 `gorm:"column:site_sub_type;type:tinyint(4);default:0" json:"site_sub_type"`
	SiteRegion       string `gorm:"column:site_region;type:varchar(255)" json:"site_region"`
	SiteCity         string `gorm:"column:site_city;type:varchar(255)" json:"site_city"`
	SiteState        string `gorm:"column:site_state;type:varchar(255)" json:"site_state"`
	SiteDistrict     string `gorm:"column:site_district;type:varchar(255)" json:"site_district"`
	SiteStreet       string `gorm:"column:site_street;type:varchar(255)" json:"site_street"`
	DetailAddress    string `gorm:"column:detail_address;type:varchar(255)" json:"detail_address"`
	Longtitude       string `gorm:"column:longtitude;type:varchar(255)" json:"longtitude"`
	Lantitude        string `gorm:"column:lantitude;type:varchar(255)" json:"lantitude"`
	Name             string `gorm:"column:name;type:varchar(255)" json:"name"`
	ZipCode          string `gorm:"column:zip_code;type:varchar(255)" json:"zip_code"`
	Phone            string `gorm:"column:phone;type:varchar(255)" json:"phone"`
	UpdateCancelRule uint64 `gorm:"column:update_cancel_rule;type:tinyint(4);default:0" json:"update_cancel_rule"`
	CbType           uint64 `gorm:"column:cb_type;type:tinyint(4);default:0" json:"cb_type"`
	CreateTime       uint64 `gorm:"autoCreateTime;column:create_time;type:int(11);NOT NULL" json:"create_time"`
	UpdateTime       uint64 `gorm:"autoUpdateTime;column:update_time;type:int(11);NOT NULL" json:"update_time"`
}

func (m *SiteResourceTab) TableName() string {
	return "site_resource_tab"
}
