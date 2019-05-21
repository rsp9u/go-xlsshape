package shape

import (
	"encoding/xml"
)

// XForm specifies the transform to be applied to the corresponding graphic frame
type XForm struct {
	XMLName xml.Name `xml:"a:xfrm"`
	FlipH   string   `xml:"flipH,attr,omitempty"`
	FlipV   string   `xml:"flipV,attr,omitempty"`
}

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

// LineProperties is a line property around the shape.
type LineProperties struct {
	XMLName xml.Name    `xml:"a:ln"`
	Fill    *SolidFill  `xml:",omitempty"`
	Dash    *PresetDash `xml:",omitempty"`
	Head    *LineEnd    `xml:"a:headEnd,omitempty"`
	Tail    *LineEnd    `xml:"a:tailEnd,omitempty"`
}

// PresetDash is a preset type for line dash.
type PresetDash struct {
	XMLName xml.Name `xml:"a:prstDash"`
	Value   string   `xml:"val,attr"`
}

// LineEnd is a type of terminal shape of the line.
type LineEnd struct {
	Type string `xml:"type,attr"`
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

// TextParticularProperties is a set of the run particular properties for the text.
type TextParticularProperties struct {
	XMLName xml.Name `xml:"a:pPr"`
	Align   string   `xml:"algn,attr"`
}

// TextRunProperties is a set of the run level properties for the text.
type TextRunProperties struct {
	XMLName  xml.Name `xml:"a:rPr"`
	Kumimoji string   `xml:"kumimoji,attr"`
	Lang     string   `xml:"lang,attr"`
	AltLang  string   `xml:"altLang,attr"`
	Size     string   `xml:"sz,attr"`
}
