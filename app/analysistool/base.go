package analysistool

import "strings"

const (
	Data1   = 1
	Data30  = 30
	Data90  = 90
	Data180 = 180
	Data365 = 365
)

type Handle struct {
	StockId string `json:"Id"`

	// 三大法人買賣統計
	ThreefoundationDataDay1   map[string]*Threefoundation `json:"threefoundationData1"`   // 每日 <20240101,data>,<20240102,data>,<20240103,data>...
	ThreefoundationDataDay30  map[string]*Threefoundation `json:"threefoundationData30"`  // 每月 <01,[]data>,<02,[]data>,<03,[]data>...
	ThreefoundationDataDay90  map[string]*Threefoundation `json:"threefoundationData90"`  // 每季 <Q1,[]data>,<Q2,[]data>,<Q3,[]data>,<Q4,[]data>
	ThreefoundationDataDay180 map[string]*Threefoundation `json:"threefoundationData180"` // 每半年 <H1,[]data>,<H2,[]data>
	ThreefoundationDataDay365 map[string]*Threefoundation `json:"threefoundationData365"` // 每年 <2024,[]data>, <2025,[]data>

	// 收盤價資料統計
	PriceDataDay1   map[string]*Price `json:"priceData1"`   // 每日 <20240101,data>,<20240102,data>,<20240103,data>...
	PriceDataDay30  map[string]*Price `json:"priceData30"`  // 每月 <01,data>,<02,data>,<03,data>...
	PriceDataDay90  map[string]*Price `json:"priceData90"`  // 每季 <Q1,data>,<Q2,data>,<Q3,data>,<Q4,data>
	PriceDataDay180 map[string]*Price `json:"priceData180"` // 每半年 <H1,data>,<H2,data>
	PriceDataDay365 map[string]*Price `json:"priceData365"` // 每年 <2024,data>, <2025,data>
}

// 收盤資料
type Price struct {
	// Company_id         string  `json:"company_id" gorm:"column:company_id"`                 // 個股編號
	// Company_name       string  `json:"company_name" gorm:"column:company_name"`             // 個股名稱
	Transaction_number int     `json:"transaction_number" gorm:"column:transaction_number"` // 交易股數/股
	Transaction_count  int     `json:"transaction_count" gorm:"column:transaction_count"`   // 成交筆數
	Transaction_amount int     `json:"transaction_amount" gorm:"column:transaction_amount"` // 成交金額
	Price_open         float32 `json:"price_open" gorm:"column:price_open"`                 // 開盤價
	Price_close        float32 `json:"price_close" gorm:"column:price_close"`               // 收盤價
	Price_max          float32 `json:"price_max" gorm:"column:price_max"`                   // 最高價
	Price_min          float32 `json:"price_min" gorm:"column:price_min"`                   // 最低價
}

// 三大法人資料
type Threefoundation struct {
	Global_china_stockbanker_buy  int `json:"global_china_stockbanker_buy" gorm:"column:global_china_stockbanker_buy"`   // 外陸資自營商買
	Global_china_stockbanker_sell int `json:"global_china_stockbanker_sell" gorm:"column:global_china_stockbanker_sell"` // 外陸資自營商賣
	Global_china_stockbanker_diff int `json:"global_china_stockbanker_diff" gorm:"column:global_china_stockbanker_diff"` // 外陸資自營商差
	Global_stockbanker_buy        int `json:"global_stockbanker_buy" gorm:"column:global_stockbanker_buy"`               // 外資自營商買
	Global_stockbanker_sale       int `json:"global_stockbanker_sale" gorm:"column:global_stockbanker_sale"`             // 外資自營商賣
	Global_stockbanker_diff       int `json:"global_stockbanker_diff" gorm:"column:global_stockbanker_diff"`             // 外資自營商差
	Stock_foundation_buy          int `json:"stock_foundation_buy" gorm:"column:stock_foundation_buy"`                   // 投資信託基金買
	Stock_foundation_sell         int `json:"stock_foundation_sell" gorm:"column:stock_foundation_sell"`                 // 投資信託基金賣
	Stockbanker_self_buy          int `json:"stockbanker_self_buy" gorm:"column:stockbanker_self_buy"`                   // 自營商自行買
	Stockbanker_self_sell         int `json:"stockbanker_self_sell" gorm:"column:stockbanker_self_sell"`                 // 自營商自行賣
	Stockbanker_self_diff         int `json:"stockbanker_self_diff" gorm:"column:stockbanker_self_diff"`                 // 自營商自行買賣差
	Stockbanker_hedging_buy       int `json:"stockbanker_hedging_buy" gorm:"column:stockbanker_hedging_buy"`             // 自營商避險買
	Stockbanker_hedging_sell      int `json:"stockbanker_hedging_sell" gorm:"column:stockbanker_hedging_sell"`           // 自營商避險賣
}

func New(stockId string) *Handle {
	return &Handle{
		StockId:                   stockId,
		ThreefoundationDataDay1:   make(map[string]*Threefoundation),
		ThreefoundationDataDay30:  make(map[string]*Threefoundation),
		ThreefoundationDataDay90:  make(map[string]*Threefoundation),
		ThreefoundationDataDay180: make(map[string]*Threefoundation),
		ThreefoundationDataDay365: make(map[string]*Threefoundation),
		PriceDataDay1:             make(map[string]*Price),
		PriceDataDay30:            make(map[string]*Price),
		PriceDataDay90:            make(map[string]*Price),
		PriceDataDay180:           make(map[string]*Price),
		PriceDataDay365:           make(map[string]*Price),
	}
}

// 讀取標的分析暫存資料
func (self *Handle) Load() {

}

// 更新股票收盤價資料
//
// @params date [2024-01-01]
//
// @params data dao.Company_stock
func (self *Handle) UpdateCloseData(date string, data *Price) {
	key_day1, key_day30, key_day90, key_day180, key_day365 := timeTool(date)
	self.PriceDataDay1[key_day1] = data

	if _, ok := self.PriceDataDay30[key_day30]; !ok {
		self.PriceDataDay30[key_day30] = &Price{}
	}
	dataSum(self.PriceDataDay30[key_day30], data)

	if _, ok := self.PriceDataDay90[key_day90]; !ok {
		self.PriceDataDay90[key_day90] = &Price{}
	}
	dataSum(self.PriceDataDay90[key_day90], data)

	if _, ok := self.PriceDataDay180[key_day180]; !ok {
		self.PriceDataDay180[key_day180] = &Price{}
	}
	dataSum(self.PriceDataDay180[key_day180], data)

	if _, ok := self.PriceDataDay365[key_day365]; !ok {
		self.PriceDataDay365[key_day365] = &Price{}
	}
	dataSum(self.PriceDataDay365[key_day365], data)

}

func timeTool(date string) (key_day1, key_day30, key_day90, key_day180, key_day365 string) {
	timeSplice := strings.Split(date, "-")
	key_day1 = strings.Join(timeSplice, "")
	key_day30 = timeSplice[1]
	key_day365 = timeSplice[0]

	switch timeSplice[1] {
	case "01", "02", "03":
		key_day90 = "Q1"
		key_day180 = "H1"
	case "04", "05", "06":
		key_day90 = "Q2"
		key_day180 = "H1"
	case "07", "08", "09":
		key_day90 = "Q3"
		key_day180 = "H2"
	case "10", "11", "12":
		key_day90 = "Q4"
		key_day180 = "H2"
	}

	return
}

// 將 source 資料累計到目標
//
// @params target 目標
func dataSum(target, source *Price) {
	target.Transaction_number += source.Transaction_number
	target.Transaction_count += source.Transaction_count
	target.Transaction_amount += source.Transaction_amount
	target.Price_open += source.Price_open
	target.Price_close += source.Price_close
	target.Price_max += source.Price_max
	target.Price_min += source.Price_min
}
