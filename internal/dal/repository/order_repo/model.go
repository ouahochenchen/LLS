package order_repo

type LineOrderTab struct {
	OrderId       uint    `gorm:"column:order_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"order_id"`
	BuyerName     string  `gorm:"column:buyer_name;type:varchar(255)" json:"buyer_name"`
	BuyerAddress  string  `gorm:"column:buyer_address;type:varchar(255)" json:"buyer_address"`
	BuyerPhone    string  `gorm:"column:buyer_phone;type:varchar(255)" json:"buyer_phone"`
	GoodsType     int     `gorm:"column:goods_type;type:int(11);NOT NULL" json:"goods_type"`
	SellerName    string  `gorm:"column:seller_name;type:varchar(255)" json:"seller_name"`
	SellerAddress string  `gorm:"column:seller_address;type:varchar(255)" json:"seller_address"`
	SellerPhone   string  `gorm:"column:seller_phone;type:varchar(255)" json:"seller_phone"`
	PackageHeight int     `gorm:"column:package_height;type:int(11);default:0" json:"package_height"`
	PackageWeight int     `gorm:"column:package_weight;type:int(11);default:0" json:"package_weight"`
	Price         float64 `gorm:"column:price;type:decimal(10,4)" json:"price"`
	OrderStatus   int     `gorm:"column:order_status;type:tinyint(4);NOT NULL" json:"order_status"`
	LineId        int     `gorm:"column:line_id;type:int(11);NOT NULL" json:"line_id"`
	CreateTime    int     `gorm:"column:create_time;type:int(11);NOT NULL" json:"create_time"`
	UpdateTime    int     `gorm:"column:update_time;type:int(11);NOT NULL" json:"update_time"`
}

func (m *LineOrderTab) TableName() string {
	return "line_order_tab"
}
