package app

import (
	"bk_analysis/app/analysistool"
	arango "bk_analysis/arangodb"
	"bk_analysis/config"
	"bk_analysis/dao"
	"bk_analysis/service/dba"
	"time"
)

type Handle struct {
	dba     dba.IHandle
	dataMap map[string]*analysistool.Handle
}

func New(setting config.Env) *Handle {
	return &Handle{
		dba:     dba.New(setting),
		dataMap: make(map[string]*analysistool.Handle),
	}
}

func (self *Handle) AddStocId(stockId string) {
	self.dataMap[stockId] = analysistool.New(stockId)
}

func (self *Handle) RunCloseData() {
	// self.dba.ReadStockClosePrice()

	keys := []string{}
	st := time.Date(2024, 04, 01, 0, 0, 0, 0, time.Local)
	et := time.Date(2024, 04, 30, 0, 0, 0, 0, time.Local)

	for st.Before(et) {
		keys = append(keys, st.Format("2006-01-02"))
		st = st.AddDate(0, 0, 1)
	}

	data_Price := arango.GetPrice(keys)

	// mockDatas := []dao.Company_stock{}

	var analyTool *analysistool.Handle
	var ok bool
	for key_date, row_Price := range data_Price {
		for _, data := range row_Price {
			if analyTool, ok = self.dataMap[data.Company_id]; ok {
				priceData := toPrice(&data)
				analyTool.UpdateCloseData(key_date, priceData)
			}
		}
	}
}

func toPrice(data *dao.Company_stock) *analysistool.Price {
	res := &analysistool.Price{}
	res.Transaction_number = data.Transaction_number
	res.Transaction_count = data.Transaction_count
	res.Transaction_amount = data.Transaction_amount
	res.Price_open = data.Price_open
	res.Price_close = data.Price_close
	res.Price_max = data.Price_max
	res.Price_min = data.Price_min
	return res
}
