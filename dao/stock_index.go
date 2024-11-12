package dao

// 大盤指數
type Stock_index struct {
	Date         string  `json:"date" gorm:"column:date"`
	Index_name   string  `json:"index_name" gorm:"column:index_name"`     // 指數名稱
	Index_number float32 `json:"index_number" gorm:"column:index_number"` // 收盤指數
}

func (self *Stock_index) TableName() string {
	return "stock_index"
}
