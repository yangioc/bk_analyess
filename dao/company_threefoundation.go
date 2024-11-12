package dao

// 個股每日三大法人買賣
type Company_threefoundation struct {
	Date                          string `json:"date" gorm:"column:date"`
	Company_id                    string `json:"company_id" gorm:"column:company_id"`                                       // 個股編號
	Company_name                  string `json:"company_name" gorm:"column:company_name"`                                   // 個股名稱
	Global_china_stockbanker_buy  int    `json:"global_china_stockbanker_buy" gorm:"column:global_china_stockbanker_buy"`   // 外陸資自營商買
	Global_china_stockbanker_sell int    `json:"global_china_stockbanker_sell" gorm:"column:global_china_stockbanker_sell"` // 外陸資自營商賣
	Global_china_stockbanker_diff int    `json:"global_china_stockbanker_diff" gorm:"column:global_china_stockbanker_diff"` // 外陸資自營商差
	Global_stockbanker_buy        int    `json:"global_stockbanker_buy" gorm:"column:global_stockbanker_buy"`               // 外資自營商買
	Global_stockbanker_sale       int    `json:"global_stockbanker_sale" gorm:"column:global_stockbanker_sale"`             // 外資自營商賣
	Global_stockbanker_diff       int    `json:"global_stockbanker_diff" gorm:"column:global_stockbanker_diff"`             // 外資自營商差
	Stock_foundation_buy          int    `json:"stock_foundation_buy" gorm:"column:stock_foundation_buy"`                   // 投資信託基金買
	Stock_foundation_sell         int    `json:"stock_foundation_sell" gorm:"column:stock_foundation_sell"`                 // 投資信託基金賣
	Stock_foundation_diff         int    `json:"stock_foundation_diff" gorm:"column:stock_foundation_diff"`                 // 投資信託基金差
	Stockbanker_self_buy          int    `json:"stockbanker_self_buy" gorm:"column:stockbanker_self_buy"`                   // 自營商自行買
	Stockbanker_self_sell         int    `json:"stockbanker_self_sell" gorm:"column:stockbanker_self_sell"`                 // 自營商自行賣
	Stockbanker_self_diff         int    `json:"stockbanker_self_diff" gorm:"column:stockbanker_self_diff"`                 // 自營商自行買賣差
	Stockbanker_hedging_buy       int    `json:"stockbanker_hedging_buy" gorm:"column:stockbanker_hedging_buy"`             // 自營商避險買
	Stockbanker_hedging_sell      int    `json:"stockbanker_hedging_sell" gorm:"column:stockbanker_hedging_sell"`           // 自營商避險賣
	Stockbanker_hedging_diff      int    `json:"stockbanker_hedging_diff" gorm:"column:stockbanker_hedging_diff"`           // 自營商避險差
	Stockbanker_diff              int    `json:"stockbanker_diff" gorm:"column:stockbanker_diff"`                           // 自營商總差
	Total_diff                    int    `json:"total_diff" gorm:"column:total_diff"`                                       // 三大法人總差
}

func (self *Company_threefoundation) TableName() string {
	return "company_threefoundation"
}
