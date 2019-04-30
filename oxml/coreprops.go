package oxml

import (
	"encoding/xml"
	"log"
)

// CoreProps is 'Core Properties' that is defined in OPC.
type CoreProps struct {
	XMLName           xml.Name `xml:"cp:coreProperties"`
	PropNamespace     string   `xml:"xmlns:cp,attr"`
	DcNameSpace       string   `xml:"xmlns:dc,attr"`
	DcTermsNameSpace  string   `xml:"xmlns:dcterms,attr"`
	DcmiTypeNameSpace string   `xml:"xmlns:dcmitype,attr"`
	XSINameSpace      string   `xml:"xmlns:xsi,attr"`
}

// NewCoreProps creates new core file properties.
func NewCoreProps() *CoreProps {
	return &CoreProps{
		PropNamespace:     xmlnsCoreProperties,
		DcNameSpace:       xmlnsDc,
		DcTermsNameSpace:  xmlnsDcTerms,
		DcmiTypeNameSpace: xmlnsDcmiType,
		XSINameSpace:      xmlnsXSI,
	}
}

// Path returns the file path in the archive.
func (cp *CoreProps) Path() string {
	return "docProps/core.xml"
}

// Content returns an xml string generated from object contents.
func (cp *CoreProps) Content() string {
	content, err := DefaultEncode(cp)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
