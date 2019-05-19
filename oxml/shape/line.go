package shape

import (
	"encoding/xml"
)

// Line is a kind of shapes.
type Line struct {
	startX, startY int
	endX, endY     int
	dashType       string
	headType       string
	tailType       string
	color          string
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

// SetStartPos sets the position of the start of this.
func (ln *Line) SetStartPos(x, y int) {
	ln.startX = x
	ln.startY = y
}

// SetEndPos sets the position of the end of this.
func (ln *Line) SetEndPos(x, y int) {
	ln.endX = x
	ln.endY = y
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
		From: NewCellAnchorFrom(ln.startX, ln.startY),
		To:   NewCellAnchorTo(ln.endX, ln.endY),
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
