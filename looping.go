package main

import 	("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SitemapIndex struct {
	// Locations is the value which is an array of the location type
	// Location is in xml under sitemap tag (for unmarshalling)
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	// bytes is what we unmarshal
	// we unmarshal to memory of s
	xml.Unmarshal(bytes, &s)
	
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}

}
