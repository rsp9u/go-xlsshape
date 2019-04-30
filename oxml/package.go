package oxml

import (
	"archive/zip"
	"bytes"
	"log"
)

// Package implements OpenXML SpreadSheetPackage.
type Package struct {
	parts []Part
}

// Add adds the given part into this package.
func (p *Package) Add(part Part) {
	p.parts = append(p.parts, part)
}

// List lists the all filepaths in this package.
func (p *Package) List() []string {
	paths := []string{}
	for _, part := range p.parts {
		paths = append(paths, part.Path())
	}
	return paths
}

// Packaging creates a new Excel file from the given files.
func (p *Package) Packaging() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	defer w.Close()

	for _, part := range p.parts {
		f, err := w.Create(part.Path())
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(part.Content()))
		if err != nil {
			log.Fatal(err)
		}
	}

	return buf, nil
}
