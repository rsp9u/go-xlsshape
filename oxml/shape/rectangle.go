package shape

import (
	"encoding/xml"
	"strconv"
)

// Rectangle is a kind of shapes.
type Rectangle struct {
	left, top     int
	width, height int
	text, lang    string
	fillColor     string
	lineColor     string
	noFill        bool
	noLine        bool
	geoType       string
	wrapType      string
	fontSize      int
	hAlign        string
	vAlign        string
}

// NewRectangle creates a rectangle.
func NewRectangle() *Rectangle {
	return &Rectangle{
		lang:      "en-US",
		fillColor: "FFFFFF",
		lineColor: "000000",
		noFill:    false,
		noLine:    false,
		geoType:   "rect",
		wrapType:  "none",
		fontSize:  1100,
		hAlign:    "l",
		vAlign:    "t",
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

// SetNoFill sets the no-fill flag.
func (r *Rectangle) SetNoFill(f bool) {
	r.noFill = f
}

// SetNoLine sets the no-line flag.
func (r *Rectangle) SetNoLine(f bool) {
	r.noLine = f
}

// SetGeoType sets the type of geometory.
func (r *Rectangle) SetGeoType(t string) {
	r.geoType = t
}

// SetWrapType sets the type of text wrapping.
func (r *Rectangle) SetWrapType(t string) {
	r.wrapType = t
}

// SetFontSize sets the text font size with one-hundredth of the given numeric value.
func (r *Rectangle) SetFontSize(size int) {
	r.fontSize = size
}

// SetHAlign sets the horizontal alignment of text.
//
// If align == "l", aligns the text to the left
// If align == "r", aligns the text to the right
// If align == "ctr", centers the text
func (r *Rectangle) SetHAlign(align string) {
	r.hAlign = align
}

// SetVAlign sets the vertical alignment of text.
//
// If align == "t", aligns the text to the top
// If align == "b", aligns the text to the bottom
// If align == "ctr", centers the text
func (r *Rectangle) SetVAlign(align string) {
	r.vAlign = align
}

// MarshalXML generates the xml element from this and puts it to the encoder.
func (r *Rectangle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var fill, linefill *SolidFill
	if !r.noFill {
		fill = &SolidFill{Color: &RgbColor{Value: r.fillColor}}
	}
	if !r.noLine {
		linefill = &SolidFill{Color: &RgbColor{Value: r.lineColor}}
	}
	xr := struct {
		From       *CellAnchorFrom
		To         *CellAnchorTo
		Shape      XdrShape
		ClientData string `xml:"xdr:clientData"`
	}{
		From: NewCellAnchorFrom(r.left, r.top),
		To:   NewCellAnchorTo(r.left+r.width, r.top+r.height),
		Shape: XdrShape{
			NvProperties: &XdrNonVisualShapeProperties{
				Properties: &XdrNonVisualProperties{ID: "1"},
			},
			Properties: &XdrShapeProperties{
				PresetGeom: &Geom{Preset: r.geoType},
				Fill:       fill,
				Line:       &LineProperties{Fill: linefill},
			},
			TextBody: &XdrTextBody{
				Properties: &TextBodyProperties{
					VerticalOverflow:   "clip",
					HorizontalOverflow: "clip",
					Wrap:               r.wrapType,
					RtlCol:             "0",
					Anchor:             r.vAlign,
				},
				PProperties: &TextParticularProperties{
					Align: r.hAlign,
				},
				RProperties: &TextRunProperties{
					Kumimoji: "1",
					Lang:     r.lang,
					AltLang:  "en-US",
					Size:     strconv.Itoa(r.fontSize),
				},
				Text: r.text,
			},
		},
	}
	return e.EncodeElement(xr, xml.StartElement{Name: xml.Name{Local: "xdr:twoCellAnchor"}})
}
