package line_repo

type LineResourceTab struct {
	LineId          uint64 `gorm:"column:line_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"line_id"`
	LineType        int    `gorm:"column:line_type;type:tinyint(4);default:0" json:"line_type"`
	LogisticsServer string `gorm:"column:logistics_server;type:varchar(255)" json:"logistics_server"`
	LineRegion      string `gorm:"column:line_region;type:varchar(255)" json:"line_region"`
	PickupCountry   string `gorm:"column:pickup_country;type:varchar(255)" json:"pickup_country"`
	DeliverCountry  string `gorm:"column:deliver_country;type:varchar(255)" json:"deliver_country"`
	PickupMode      int    `gorm:"column:pickup_mode;type:tinyint(4)" json:"pickup_mode"`
	DeliverMode     int    `gorm:"column:deliver_mode;type:tinyint(4)" json:"deliver_mode"`
	IsCode          int    `gorm:"column:is_code;type:tinyint(4)" json:"is_code"`
	CreateTime      int    `gorm:"column:create_time;type:int(11);NOT NULL" json:"create_time"`
	UpdateTime      int    `gorm:"column:update_time;type:int(11);NOT NULL" json:"update_time"`
}

func (m *LineResourceTab) TableName() string {
	return "line_resource_tab"
}
