package oxml

import (
	"log"
	"os"
)

// Spreadsheet creates a consumable Excel file.
type Spreadsheet struct {
	pkg          *Package
	contenttypes *ContentTypes
	workbook     *Workbook
	worksheets   []*Worksheet
	coreprops    *CoreProps
	appprops     *AppProps
}

// NewSpreadsheet creates a new spread sheet.
func NewSpreadsheet() *Spreadsheet {
	ct := NewContentTypes()
	ct.AddDefault(DefaultType{Extension: "rels", ContentType: "application/vnd.openxmlformats-package.relationships+xml"})
	ct.AddDefault(DefaultType{Extension: "xml", ContentType: "application/xml"})
	ct.AddOverride(OverrideType{PartName: "/xl/workbook.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"})
	ct.AddOverride(OverrideType{PartName: "/xl/worksheets/sheet1.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"})
	ct.AddOverride(OverrideType{PartName: "/xl/theme/theme1.xml", ContentType: "application/vnd.openxmlformats-officedocument.theme+xml"})
	ct.AddOverride(OverrideType{PartName: "/xl/styles.xml", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"})
	ct.AddOverride(OverrideType{PartName: "/docProps/core.xml", ContentType: "application/vnd.openxmlformats-package.core-properties+xml"})
	ct.AddOverride(OverrideType{PartName: "/docProps/app.xml", ContentType: "application/vnd.openxmlformats-officedocument.extended-properties+xml"})

	coreProps := NewCoreProps()
	appProps := NewAppProps()

	wb := NewWorkbook("xl/workbook.xml")
	ws := []*Worksheet{NewWorksheet("xl/worksheets/sheet1.xml")}
	wb.Add("Sheet1", "1", ws[0])

	rel := NewRelationships("_rels/.rels")
	rel.Add(Relationship{ID: "rId1", Type: typeRelationshipsDocument, Target: wb.Path()})
	rel.Add(Relationship{ID: "rId2", Type: typeRelationshipsCoreProperties, Target: coreProps.Path()})
	rel.Add(Relationship{ID: "rId3", Type: typeRelationshipsExtentedProperties, Target: appProps.Path()})

	p := Package{}
	p.Add(ct)
	p.Add(rel)
	p.Add(coreProps)
	p.Add(appProps)
	p.Add(wb)
	p.Add(wb.Relationships())
	for _, sheet := range ws {
		p.Add(sheet)
	}

	return &Spreadsheet{&p, ct, wb, ws, coreProps, appProps}
}

// List lists the all paths in this file.
func (s *Spreadsheet) List() []string {
	return s.pkg.List()
}

// Dump writes out the contents into the new Excel file
func (s *Spreadsheet) Dump(filename string) {
	buf, _ := s.pkg.Packaging()
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(buf.Bytes())
}
