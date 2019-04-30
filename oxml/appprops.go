package oxml

import (
	"encoding/xml"
	"log"
)

// AppProps is a on of file properties for Excel.
type AppProps struct {
	XMLName     xml.Name `xml:"Properties"`
	Namespace   string   `xml:"xmlns,attr"`
	VtNameSpace string   `xml:"xmlns:vt,attr"`
	Application string
}

// NewAppProps creates new core file properties.
func NewAppProps() *AppProps {
	return &AppProps{
		Namespace:   xmlnsExtendedProperties,
		VtNameSpace: xmlnsVTypes,
		Application: "Microsoft Excel",
	}
}

// Path returns the file path in the archive.
func (ap *AppProps) Path() string {
	return "docProps/app.xml"
}

// Content returns an xml string generated from object contents.
func (ap *AppProps) Content() string {
	content, err := DefaultEncode(ap)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
