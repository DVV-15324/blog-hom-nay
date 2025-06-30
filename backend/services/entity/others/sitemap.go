package others

import (
	"encoding/xml"
)

// Url struct định nghĩa 1 URL trong sitemap
type Url struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

// Urlset struct theo chuẩn sitemap
type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}
