package shape

import (
	"encoding/xml"
)

// XdrShape is struct.
type XdrShape struct {
	XMLName      xml.Name                     `xml:"xdr:sp"`
	NvProperties *XdrNonVisualShapeProperties `xml:",omitempty"`
	Properties   *XdrShapeProperties          `xml:",omitempty"`
	TextBody     *XdrTextBody                 `xml:",omitempty"`
}

// XdrNonVisualShapeProperties is struct.
type XdrNonVisualShapeProperties struct {
	XMLName         xml.Name                `xml:"xdr:nvSpPr"`
	Properties      *XdrNonVisualProperties `xml:",omitempty"`
	ShapeProperties string                  `xml:"xdr:cNvSpPr"`
}

// XdrNonVisualProperties is struct.
type XdrNonVisualProperties struct {
	XMLName xml.Name `xml:"xdr:cNvPr"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// XdrShapeProperties is struct.
type XdrShapeProperties struct {
	XMLName    xml.Name        `xml:"xdr:spPr"`
	XForm      *XForm          `xml:",omitempty"`
	PresetGeom *Geom           `xml:",omitempty"`
	Fill       *SolidFill      `xml:",omitempty"`
	NoFill     *NoFill         `xml:",omitempty"`
	Line       *LineProperties `xml:",omitempty"`
}

// XdrTextBody is struct.
type XdrTextBody struct {
	XMLName     xml.Name                  `xml:"xdr:txBody"`
	Properties  *TextBodyProperties       `xml:",omitempty"`
	ListStyle   string                    `xml:"a:lstStyle"`
	PProperties *TextParticularProperties `xml:"a:p>a:pPr"`
	RProperties *TextRunProperties        `xml:"a:p>a:r>a:rPr"`
	Text        string                    `xml:"a:p>a:r>a:t"`
}
