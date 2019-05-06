package oxml

import (
	"encoding/xml"
	"log"
	"strconv"
)

// Worksheet is a worksheet.
type Worksheet struct {
	XMLName       xml.Name `xml:"worksheet"`
	Namespace     string   `xml:"xmlns,attr"`
	RelNameSpace  string   `xml:"xmlns:r,attr"`
	SheetFormat   SheetFormat
	SheetData     SheetData
	DrawingRels   []DrawingRel
	path          string
	relationships *Relationships
}

// SheetFormat is a struct.
type SheetFormat struct {
	XMLName            xml.Name `xml:"sheetFormatPr"`
	DefaultColumnWidth string   `xml:"defaultColWidth,attr,omitempty"`
	DefaultRowHeight   string   `xml:"defaultRowHeight,attr,omitempty"`
	CustomHeight       string   `xml:"customHeight,attr,omitempty"`
}

// SheetData is a struct.
type SheetData struct {
	XMLName xml.Name `xml:"sheetData"`
}

// DrawingRel is a relationships of the drawing.
type DrawingRel struct {
	XMLName xml.Name `xml:"drawing"`
	ID      string   `xml:"r:id,attr"`
}

// NewWorksheet creates new worksheet.
func NewWorksheet(path string) *Worksheet {
	relpath := RelationshipPath(path)
	relationships := NewRelationships(relpath)
	return &Worksheet{
		Namespace:     xmlnsSpreadSheetMain,
		RelNameSpace:  xmlnsOfficeRelationships,
		path:          path,
		relationships: relationships,
	}
}

// Path returns the file path in the archive.
func (ws *Worksheet) Path() string {
	return ws.path
}

// Content returns an xml string generated from object contents.
func (ws *Worksheet) Content() string {
	content, err := DefaultEncode(ws)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// SetDefaultCellSize sets the default size of cell.
func (ws *Worksheet) SetDefaultCellSize(w, h string) {
	ws.SheetFormat.DefaultColumnWidth = w
	ws.SheetFormat.DefaultRowHeight = h
	ws.SheetFormat.CustomHeight = "1"
}

// AddDrawing adds a drawing relationship into this.
func (ws *Worksheet) AddDrawing(target Part) {
	rid := "rId" + strconv.Itoa(len(ws.DrawingRels)+1)
	ws.DrawingRels = append(ws.DrawingRels, DrawingRel{ID: rid})
	ws.relationships.Add(Relationship{ID: rid, Type: typeRelationshipsDrawing, Target: TargetPath(ws, target)})
}

// Relationships returns the relationships associated to this.
func (ws *Worksheet) Relationships() *Relationships {
	return ws.relationships
}
