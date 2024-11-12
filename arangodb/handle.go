package arango

import (
	"bk_analysis/dao"
	"context"
	"fmt"
	"sort"
)

func GetPrice(keys []string) map[string][]dao.Company_stock {

	payload := make([]PriceData, len(keys))

	err := _instans.Reads(context.TODO(), "stockcloseprice", keys, payload)
	fmt.Println(err)
	res := map[string][]dao.Company_stock{}

	sort.Slice(payload, func(i, j int) bool {
		return payload[i].Key < payload[j].Key
	})
	for _, data := range payload {
		if data.Key != "" {
			res[data.Key] = data.Datas
		}
	}
	return res
}

type PriceData struct {
	Key   string              `json:"_key"`
	Datas []dao.Company_stock `json:"datas"`
}
