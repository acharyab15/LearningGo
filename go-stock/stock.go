package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type PriceAggPage struct {
	Title  string
	Prices map[string]interface{}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=MSFT&apikey=demo")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s interface{}
	json.Unmarshal(bytes, &s)
	resp.Body.Close()
	m := s.(map[string]interface{})
	timeSeriesMap := m["Time Series (Daily)"]
	v := timeSeriesMap.(map[string]interface{})
	// get today's date
	today := time.Now().Local().Format("2006-01-02")
	dailyMap := v[today]
	details := dailyMap.(map[string]interface{})
	price_map := make(map[string]interface{})
	price_map["open"] = details["1. open"]
	price_map["high"] = details["2. high"]
	price_map["low"] = details["3. low"]
	price_map["close"] = details["4. close"]
	price_map["volume"] = details["5. volume"]
	p := PriceAggPage{Title: "Amazing Stock Aggregator", Prices: price_map}
	t, _ := template.ParseFiles("basic.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
