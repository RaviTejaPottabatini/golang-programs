package main

import (
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	//"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

func main() {

    

	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	//bytes, _ := ioutil.ReadAll(resp.Body)
    bytes := washPostXML
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}