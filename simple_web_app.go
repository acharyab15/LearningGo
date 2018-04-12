package main

import ("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template")

type SitemapIndex struct {
	// Locations is the value which is a slice of the string type
	// of Locations which is in xml under sitemap tag (for unmarshalling)
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// bytes is what we unmarshal
	// we unmarshal to memory of s
	xml.Unmarshal(bytes, &s)
	
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
