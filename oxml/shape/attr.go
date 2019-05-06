package shape

import (
	"encoding/xml"
)

// RgbColor is a struct to be specified a color by rgb hex values like "FFFF00".
type RgbColor struct {
	XMLName xml.Name `xml:"a:srgbClr"`
	Value   string   `xml:"val,attr"`
}

// Geom is geometry type with preset.
type Geom struct {
	XMLName      xml.Name `xml:"a:prstGeom"`
	Preset       string   `xml:"prst,attr"`
	AdjustValues string   `xml:"a:avLst"`
}

// SolidFill is a fill property of the shape.
type SolidFill struct {
	XMLName xml.Name  `xml:"a:solidFill"`
	Color   *RgbColor `xml:",omitempty"`
}

// NoFill is nothing of fill element.
type NoFill struct {
	XMLName xml.Name `xml:"a:noFill"`
}

// Line is a line property around the shape.
type Line struct {
	XMLName xml.Name   `xml:"a:ln"`
	Fill    *SolidFill `xml:",omitempty"`
}

// TextBodyProperties is a set of properties of the text box.
type TextBodyProperties struct {
	XMLName            xml.Name  `xml:"a:bodyPr"`
	VerticalOverflow   string    `xml:"vertOverflow,attr"`
	HorizontalOverflow string    `xml:"horzOverflow,attr"`
	Wrap               string    `xml:"wrap,attr"`
	RtlCol             string    `xml:"rtlCol,attr"`
	Anchor             string    `xml:"anchor,attr"`
	AutoFit            *struct{} `xml:"a:spAutoFit,omitempty"`
}

// TextProperties is a set of properties for the text.
type TextProperties struct {
	XMLName  xml.Name `xml:"a:rPr"`
	Kumimoji string   `xml:"kumimoji,attr"`
	Lang     string   `xml:"lang,attr"`
	AltLang  string   `xml:"altLang,attr"`
	Size     string   `xml:"sz,attr"`
}