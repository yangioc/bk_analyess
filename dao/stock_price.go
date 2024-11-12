package dao

// 大盤統計
type Stock_price struct {
	Date               string `json:"date" gorm:"column:date"`
	Stock_name         string `json:"stock_name" gorm:"column:stock_name"`                 // 統計類型名稱
	Stock_total_amount int    `json:"stock_total_amount" gorm:"column:stock_total_amount"` // 統計總成交金額
	Stock_total_number int    `json:"stock_total_number" gorm:"column:stock_total_number"` // 統計總成交股數
	Stock_total_count  int    `json:"stock_total_count" gorm:"column:stock_total_count"`   // 統計總成交比數
}

func (self *Stock_price) TableName() string {
	return "stock_price"
}
