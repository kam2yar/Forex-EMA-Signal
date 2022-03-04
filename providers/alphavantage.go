package providers

import (
	"bed/helpers"
	"strconv"
)

type Alphavantage struct{}

func (provider Alphavantage) GetSymbols() []string {
	return []string{"EURUSD", "USDJPY", "GBPUSD", "AUDUSD", "USDCAD", "USDCHF", "NZDUSD", "EURGBP", "EURAUD", "EURCHF",
		"EURJPY", "EURNZD", "GBPEUR", "GBPJPY", "GBPAUD", "GBPCAD", "GBPCHF", "GBPNZD", "CADCHF", "CADJPY"}
}

func (provider Alphavantage) GetValues(symbol string, interval string, period int) map[string]float64 {
	url := "https://www.alphavantage.co/query?function=EMA&symbol=" + symbol + "&interval=" + interval + "&time_period=" + strconv.Itoa(period) + "&series_type=close&apikey=" + helpers.Env("ALPHAVANTAGE_API_KEY")

	jsonResult := request(url)

	result := finalResult{}

	for date, item := range jsonResult["Technical Analysis: EMA"].(map[string]interface{}) {
		value := item.(map[string]interface{})["EMA"].(string)
		result[date], _ = strconv.ParseFloat(value, 64)
	}

	return result
}
