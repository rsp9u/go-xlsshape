package oxml

import (
	"encoding/xml"
	"log"
)

// ContentTypes specify the content types in the archive.
type ContentTypes struct {
	XMLName   xml.Name `xml:"Types"`
	Namespace string   `xml:"xmlns,attr"`
	Defaults  []DefaultType
	Overrides []OverrideType
}

// DefaultType is the content type per extension.
type DefaultType struct {
	XMLName     xml.Name `xml:"Default"`
	Extension   string   `xml:",attr"`
	ContentType string   `xml:",attr"`
}

// OverrideType is the content type per file.
type OverrideType struct {
	XMLName     xml.Name `xml:"Override"`
	PartName    string   `xml:",attr"`
	ContentType string   `xml:",attr"`
}

// NewContentTypes creates a new content types.
func NewContentTypes() *ContentTypes {
	return &ContentTypes{Namespace: xmlnsContentTypes}
}

// Path returns the file path in the archive.
func (c *ContentTypes) Path() string {
	return "[Content_Types].xml"
}

// Content returns an xml string generated from object contents.
func (c *ContentTypes) Content() string {
	content, err := DefaultEncode(c)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// AddDefault adds a default type.
func (c *ContentTypes) AddDefault(item DefaultType) {
	c.Defaults = append(c.Defaults, item)
}

// AddOverride adds an override type.
func (c *ContentTypes) AddOverride(item OverrideType) {
	c.Overrides = append(c.Overrides, item)
}
