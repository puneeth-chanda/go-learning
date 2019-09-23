package main

import("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

type SitemapIndex struct{
	Locations []Location `xml:"sitemap"`
}
type Location struct{
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}
func main(){
	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes,_ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	
}