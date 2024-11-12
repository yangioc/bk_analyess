package dao

// 個股日本益比、殖利率及股價淨值比
type CompanyStockAnalysis struct {
	Stat       string          `json:"stat"`
	Date       string          `json:"date"`
	Title      string          `json:"title"`
	Fields     []string        `json:"fields"`
	Data       [][]interface{} `json:"data"`
	SelectType string          `json:"selectType"`
	Total      int64           `json:"total"`
}

type Company_stock_analysis struct {
	Date                string  `json:"date" gorm:"column:date"`
	Company_id          string  `json:"company_id" gorm:"column:company_id"`                   // 個股編號
	Company_name        string  `json:"company_name" gorm:"column:company_name"`               // 個股名稱
	Yield               float32 `json:"yield" gorm:"column:yield"`                             // 殖利率/%, 殖利率 = 每股股利／收盤價*100%
	Yield_year          string  `json:"yield_year" gorm:"column:yield_year"`                   // 股利年度
	PE                  float32 `json:"pe" gorm:"column:pe"`                                   // 本益比/%, 本益比(PE) = 收盤價／每股參考稅後純益(EPS)
	PBR                 float32 `json:"pbr" gorm:"column:pbr"`                                 // 股價淨值比/%, 股價淨值比(PBR) = 收盤價／每股參考淨值
	Analysis_year_month string  `json:"analysis_year_month" gorm:"column:analysis_year_month"` // 財報 年/季
}

func (self *Company_stock_analysis) TableName() string {
	return "company_stock_analysis"
}
