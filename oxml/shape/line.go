package shape

import (
	"encoding/xml"
)

// Line is a kind of shapes.
type Line struct {
	left, top     int
	right, bottom int
	dashType      string
	headType      string
	tailType      string
	color         string
}

// NewLine creates a rectangle.
func NewLine() *Line {
	return &Line{
		dashType: "",
		headType: "",
		tailType: "",
		color:    "000000",
	}
}

// SetLeftTop sets top and left of this.
func (ln *Line) SetLeftTop(l, t int) {
	ln.left = l
	ln.top = t
}

// SetRightBottom sets bottom and right of this.
func (ln *Line) SetRightBottom(r, b int) {
	ln.right = r
	ln.bottom = b
}

// SetDashType sets the type of line dash of this.
func (ln *Line) SetDashType(t string) {
	ln.dashType = t
}

// SetHeadType sets the type of the head of this.
func (ln *Line) SetHeadType(t string) {
	ln.headType = t
}

// SetTailType sets the type of the tail of this.
func (ln *Line) SetTailType(t string) {
	ln.tailType = t
}

// SetColor sets the color of this.
func (ln *Line) SetColor(c string) {
	ln.color = c
}

// MarshalXML generates the xml element from this and puts it to the encoder.
func (ln *Line) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var (
		dash       *PresetDash
		head, tail *LineEnd
	)
	if ln.dashType != "" {
		dash = &PresetDash{Value: ln.dashType}
	}
	if ln.headType != "" {
		head = &LineEnd{Type: ln.headType}
	}
	if ln.tailType != "" {
		tail = &LineEnd{Type: ln.tailType}
	}

	xr := struct {
		From       *CellAnchorFrom
		To         *CellAnchorTo
		Shape      XdrShape
		ClientData string `xml:"xdr:clientData"`
	}{
		From: NewCellAnchorFrom(ln.left, ln.top),
		To:   NewCellAnchorTo(ln.right, ln.bottom),
		Shape: XdrShape{
			NvProperties: &XdrNonVisualShapeProperties{
				Properties: &XdrNonVisualProperties{ID: "1"},
			},
			Properties: &XdrShapeProperties{
				PresetGeom: &Geom{Preset: "straightConnector1"},
				Line: &LineProperties{
					Fill: &SolidFill{Color: &RgbColor{Value: ln.color}},
					Dash: dash,
					Head: head,
					Tail: tail,
				},
			},
		},
	}
	return e.EncodeElement(xr, xml.StartElement{Name: xml.Name{Local: "xdr:twoCellAnchor"}})
}
