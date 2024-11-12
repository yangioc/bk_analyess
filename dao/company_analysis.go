package dao

type Company_analysis struct {
	Year       int    `json:"year" gorm:"column:year"`
	Session    string `json:"session" gorm:"column:session"`
	Company_id string `json:"company_id" gorm:"column:company_id"`

	Data_type                    string  `json:"data_type" gorm:"column:data_type"`                                       // 資料類型: session:單季 session_sum:季類計 year:年
	Revenue                      float64 `json:"revenue" gorm:"column:revenue"`                                           // 總營收	公司尚未扣除任何費用的總收入。
	Gross_profit_from_operations float64 `json:"gross_profit_from_operations" gorm:"column:gross_profit_from_operations"` // 營業毛利
	Gross_margin                 float64 `json:"gross_margin" gorm:"column:gross_margin"`                                 // 毛利率	代表產品的成本與總收入的關係。	(一項產品扣掉直接成本，可從售價獲利幾%)
	Operating_income             float64 `json:"operating_income" gorm:"column:operating_income"`                         // 營業利潤	公司透過商業活動所獲得的收入，通常是提供產品及服務而來。
	Operation_margin             float64 `json:"operation_margin" gorm:"column:operation_margin"`                         // 營業利率	營業利率則是營業利潤與總營收的關係。	(一項產品扣掉所有成本，可從售價獲利幾%)
	Net_income                   float64 `json:"net_income" gorm:"column:net_income"`                                     // 淨收入	是評估公司獲利能力的關鍵。
	Earnings_per_share           float64 `json:"earnings_per_share" gorm:"column:earnings_per_share"`                     // 每股盈餘 (EPS)	是公司的獲利指標，而通常EPS與股價也會有非常大的關連性。
	Book_value_per_share         float64 `json:"book_value_per_share" gorm:"column:book_value_per_share"`                 // 每股淨資產	企業把所有資產變現，並償還債務後所剩餘的部份。
	Operation_cash_flow          float64 `json:"operation_cash_flow" gorm:"column:operation_cash_flow"`                   // 營運現金流	公司生產營運所帶來的現金流入和流出的數值，包括所有折舊與攤提。
	Cap_spending                 float64 `json:"cap_spending" gorm:"column:cap_spending"`                                 // 資本支出	購入未來可替公司增加經濟效益的固定資產。如廠房、機器設備等。
	Free_cash_flow               float64 `json:"free_cash_flow" gorm:"column:free_cash_flow"`                             // 自由現金流	是公司可自由運用的現金流，衡量公司手上持有現金的狀況。
	Return_on_equity             float64 `json:"return_on_equity" gorm:"column:return_on_equity"`                         // 股東權益報酬率	公司利用資產淨值產生利益的能力，也就是公司替股東賺錢的能力。
	Return_on_assets             float64 `json:"return_on_assets" gorm:"column:return_on_assets"`                         // 資產報酬率	由於 ROE 計算時不納入債務，因此若企業屬於高財務槓桿，意即資產中有較高比例的舉債，ROE 便會遠高於 ROA
	Interest_coverage            float64 `json:"interest_coverage" gorm:"column:interest_coverage"`                       // 利息保障倍數	衡量一間公司支付負債利息的能力，如果數值愈高就代表企業還錢的能力越好。
	Net_margin                   float64 `json:"net_margin" gorm:"column:net_margin"`                                     // 淨利率	公司從總營收中計算淨收入的比例
}

func (self *Company_analysis) TableName() string {
	return "company_analysis"
}

const (
	Company_analysis_Data_type_session     = "session"
	Company_analysis_Data_type_session_sum = "session_sum"
	Company_analysis_Data_type_year        = "year"
)
