package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const API_KEY = "MGBL4982FGJ54MOB"

type PriceAggPage struct {
	Title  string
	Prices map[string]interface{}
}

func stockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		price_map := make(map[string]interface{})
		p := PriceAggPage{Title: "Amazing Stock Aggregator", Prices: price_map}
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, p)
	} else {
		fmt.Println("method: ", r.Method)
		r.ParseForm()
		stock_name := r.Form.Get("stock_name")
		fmt.Println("Stock name: ", stock_name)
		query := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", stock_name, API_KEY)
		fmt.Println(query)
		resp, err := http.Get(query)
		if err != nil {
			fmt.Println("Error getting stock info", err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading stock response", err)
		}
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
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stock", stockHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
