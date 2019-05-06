package shape

import (
	"encoding/xml"
)

// Rectangle is a kind of shapes.
type Rectangle struct {
	left, top     int
	width, height int
	text, lang    string
	fillColor     string
	lineColor     string
}

// NewRectangle creates a rectangle.
func NewRectangle() *Rectangle {
	return &Rectangle{
		lang:      "en-US",
		fillColor: "FFFFFF",
		lineColor: "000000",
	}
}

// SetLeftTop sets top and left of this.
func (r *Rectangle) SetLeftTop(l, t int) {
	r.left = l
	r.top = t
}

// SetSize sets top and left of this.
func (r *Rectangle) SetSize(w, h int) {
	r.width = w
	r.height = h
}

// SetText sets inner text of this.
func (r *Rectangle) SetText(t, lang string) {
	r.text = t
	r.lang = lang
}

// SetFillColor sets the color used to fill this.
func (r *Rectangle) SetFillColor(c string) {
	r.fillColor = c
}

// SetLineColor sets the color of the line around this.
func (r *Rectangle) SetLineColor(c string) {
	r.lineColor = c
}

// MarshalXML generates the xml element from this and puts it to the encoder.
func (r *Rectangle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xr := struct {
		From       *CellAnchorFrom
		To         *CellAnchorTo
		Shape      XdrShape
		ClientData string `xml:"xdr:clientData"`
	}{
		From: NewCellAnchorFrom(r.left, r.top),
		To:   NewCellAnchorTo(r.left+r.width, r.top+r.height),
		Shape: XdrShape{
			NvProperties: XdrNonVisualShapeProperties{
				Properties: XdrNonVisualProperties{ID: "1"},
			},
			Properties: XdrShapeProperties{
				PresetGeom: &Geom{Preset: "rect"},
				Fill:       &SolidFill{Color: &RgbColor{Value: r.fillColor}},
				Line:       &Line{Fill: &SolidFill{Color: &RgbColor{Value: r.lineColor}}},
			},
			TextBody: &XdrTextBody{
				Properties: &TextBodyProperties{
					VerticalOverflow:   "clip",
					HorizontalOverflow: "clip",
					Wrap:               "none",
					RtlCol:             "0",
					Anchor:             "t",
				},
				RProperties: &TextProperties{
					Kumimoji: "1",
					Lang:     r.lang,
					AltLang:  "en-US",
					Size:     "1100",
				},
				Text: r.text,
			},
		},
	}
	return e.EncodeElement(xr, xml.StartElement{Name: xml.Name{Local: "xdr:twoCellAnchor"}})
}