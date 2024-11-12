package dao

type Company_financial_statements struct {
	Year         int    `json:"year" gorm:"column:year"`                 // 資料年分
	Session      string `json:"session" gorm:"column:session"`           // 資料範圍/季
	Company_id   string `json:"company_id" gorm:"column:company_id"`     // 個股編號
	Data         string `json:"data" gorm:"column:data"`                 // 財報資料
	Data_version string `json:"data_version" gorm:"column:data_version"` // 財報資料格式版本
}

func (self *Company_financial_statements) TableName() string {
	return "company_financial_statements"
}

const (
	Company_financial_statements_DataVersion = "20230905001"
)
