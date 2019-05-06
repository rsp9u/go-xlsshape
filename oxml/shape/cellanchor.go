package shape

import (
	"encoding/xml"
)

// CellAnchorFrom is an anchor struct used within oneCellAnchor or twoCellAnchor.
type CellAnchorFrom struct {
	XMLName      xml.Name `xml:"xdr:from"`
	Column       int      `xml:"xdr:col"`
	ColumnOffset int      `xml:"xdr:colOff"`
	Row          int      `xml:"xdr:row"`
	RowOffset    int      `xml:"xdr:rowOff"`
}

// CellAnchorTo is an anchor struct used within twoCellAnchor.
type CellAnchorTo struct {
	XMLName      xml.Name `xml:"xdr:to"`
	Column       int      `xml:"xdr:col"`
	ColumnOffset int      `xml:"xdr:colOff"`
	Row          int      `xml:"xdr:row"`
	RowOffset    int      `xml:"xdr:rowOff"`
}

const (
	cellSize = 20
	scale    = 10000
)

func cell(left, top int) (int, int) {
	return left / cellSize, top / cellSize
}

func cellOffset(left, top int) (int, int) {
	return left % cellSize * scale, top % cellSize * scale
}

// NewCellAnchorFrom creates CellAnchorFrom with left and top pixels.
func NewCellAnchorFrom(left, top int) *CellAnchorFrom {
	cx, cy := cell(left, top)
	cox, coy := cellOffset(left, top)

	return &CellAnchorFrom{
		Column:       cx,
		Row:          cy,
		ColumnOffset: cox,
		RowOffset:    coy,
	}
}

// NewCellAnchorTo creates CellAnchorTo with left and top pixels.
func NewCellAnchorTo(left, top int) *CellAnchorTo {
	cx, cy := cell(left, top)
	cox, coy := cellOffset(left, top)

	return &CellAnchorTo{
		Column:       cx,
		Row:          cy,
		ColumnOffset: cox,
		RowOffset:    coy,
	}
}
