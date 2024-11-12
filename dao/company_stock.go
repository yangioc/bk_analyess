package dao

// 個股每日收盤價
type Company_stock struct {
	Date               string  `json:"date" gorm:"column:date"`
	Company_id         string  `json:"company_id" gorm:"column:company_id"`                 // 個股編號
	Company_name       string  `json:"company_name" gorm:"column:company_name"`             // 個股名稱
	Transaction_number int     `json:"transaction_number" gorm:"column:transaction_number"` // 交易股數/股
	Transaction_count  int     `json:"transaction_count" gorm:"column:transaction_count"`   // 成交筆數
	Transaction_amount int     `json:"transaction_amount" gorm:"column:transaction_amount"` // 成交金額
	Price_open         float32 `json:"price_open" gorm:"column:price_open"`                 // 開盤價
	Price_close        float32 `json:"price_close" gorm:"column:price_close"`               // 收盤價
	Price_max          float32 `json:"price_max" gorm:"column:price_max"`                   // 最高價
	Price_min          float32 `json:"price_min" gorm:"column:price_min"`                   // 最低價
}

func (self *Company_stock) TableName() string {
	return "company_stock"
}
