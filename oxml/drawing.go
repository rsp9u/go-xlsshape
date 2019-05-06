package oxml

import (
	"encoding/xml"
	"log"

	"github.com/rsp9u/go-xlsshape/oxml/shape"
)

// Drawing is a drawing object.
type Drawing struct {
	XMLName         xml.Name `xml:"xdr:wsDr"`
	Namespace       string   `xml:"xmlns:xdr,attr"`
	AnchorNamespace string   `xml:"xmlns:a,attr"`
	Shapes          []shape.Shape
	path            string
}

// NewDrawing creates new drawing.
func NewDrawing(path string) *Drawing {
	return &Drawing{
		Namespace:       xmlnsSpreadsheetDrawing,
		AnchorNamespace: xmlnsDrawingMLMain,
		path:            path,
	}
}

// Path returns the file path in the archive.
func (d *Drawing) Path() string {
	return d.path
}

// Content returns an xml string generated from object contents.
func (d *Drawing) Content() string {
	content, err := DefaultEncode(d)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// AddShape adds a shape to this.
func (d *Drawing) AddShape(s shape.Shape) {
	d.Shapes = append(d.Shapes, s)
}
