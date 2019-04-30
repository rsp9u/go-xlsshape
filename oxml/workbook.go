package oxml

import (
	"encoding/xml"
	"log"
	"strconv"
)

// Workbook contains one or more worksheet.
type Workbook struct {
	XMLName       xml.Name `xml:"workbook"`
	Namespace     string   `xml:"xmlns,attr"`
	RelNameSpace  string   `xml:"xmlns:r,attr"`
	Sheets        Sheets
	path          string
	relationships *Relationships
}

// Sheets is a slice of Sheet.
type Sheets struct {
	XMLName xml.Name `xml:"sheets"`
	Items   []Sheet
}

// Sheet is a relationship of the worksheet.
type Sheet struct {
	XMLName        xml.Name `xml:"sheet"`
	Name           string   `xml:"name,attr"`
	SheetID        string   `xml:"sheetId,attr"`
	RelationshipID string   `xml:"r:id,attr"`
}

// NewWorkbook creates new workbook.
func NewWorkbook(path string) *Workbook {
	relpath := RelationshipPath(path)
	relationships := NewRelationships(relpath)
	return &Workbook{
		Namespace:     xmlnsSpreadSheetMain,
		RelNameSpace:  xmlnsOfficeRelationships,
		path:          path,
		relationships: relationships,
	}
}

// Path returns the file path in the archive.
func (wb *Workbook) Path() string {
	return wb.path
}

// Content returns an xml string generated from object contents.
func (wb *Workbook) Content() string {
	content, err := DefaultEncode(wb)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// Add adds a sheet.
func (wb *Workbook) Add(name, sheetID string, target Part) {
	rid := "rId" + strconv.Itoa(len(wb.relationships.Items)+1)
	wb.Sheets.Items = append(wb.Sheets.Items, Sheet{Name: name, SheetID: sheetID, RelationshipID: rid})
	wb.relationships.Add(Relationship{ID: rid, Type: typeRelationshipsWorkSheet, Target: TargetPath(wb, target)})
}

// Relationships returns the relationships associated to this.
func (wb *Workbook) Relationships() *Relationships {
	return wb.relationships
}
