package main

import 	("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SitemapIndex struct {
	// Locations is the value which is a slice of the string type
	// of Locations which is in xml under sitemap tag (for unmarshalling)
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keyworkds"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News

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
	}
}
