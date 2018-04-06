package main


/*
var washPostXML = []byte(`
<sitemapindex>
	<sitemap>
	      <loc>http://www.washingtonpost.com/news-blogs-technology-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-lifestyle-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-blogs-lifestyle-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-entertainment-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml</loc>
	   </sitemap>
	   <sitemap>
	      <loc>http://www.washingtonpost.com/news-goingoutguide-sitemap.xml</loc>
	   </sitemap>
</sitemapindex>`)
*/

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
	
	fmt.Println(s.Locations)
}
