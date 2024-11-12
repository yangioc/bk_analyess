package dao

// 三大法人總買賣
type Stock_threefoundation struct {
	Date             string `json:"date" gorm:"column:date"`
	Stock_three_name string `json:"stock_three_name" gorm:"column:stock_three_name"` // 統計類型名稱
	Buy              int    `json:"buy" gorm:"column:buy"`                           // 總買
	Sell             int    `json:"sell" gorm:"column:sell"`                         // 總賣
	Diff             int    `json:"diff" gorm:"column:diff"`                         // 總差
}

func (self *Stock_threefoundation) TableName() string {
	return "stock_threefoundation"
}
