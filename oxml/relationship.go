package oxml

import (
	"encoding/xml"
	"log"
	"path/filepath"
)

// Relationships contains zero or more relationship.
type Relationships struct {
	XMLName   xml.Name `xml:"Relationships"`
	Namespace string   `xml:"xmlns,attr"`
	Items     []Relationship
	path      string
}

// Relationship has a relation between a source part and a target part.
type Relationship struct {
	XMLName xml.Name `xml:"Relationship"`
	ID      string   `xml:"Id,attr"`
	Type    string   `xml:"Type,attr"`
	Target  string   `xml:"Target,attr"`
}

// NewRelationships creates new relationships.
func NewRelationships(path string) *Relationships {
	return &Relationships{Namespace: xmlnsRelationships, path: path}
}

// Path returns the file path in the archive.
func (r *Relationships) Path() string {
	return r.path
}

// Content returns an xml string generated from object contents.
func (r *Relationships) Content() string {
	content, err := DefaultEncode(r)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// Add adds a relationship.
func (r *Relationships) Add(item Relationship) {
	r.Items = append(r.Items, item)
}

// RelationshipPath returns the relationship file path that is relative to the given file.
func RelationshipPath(path string) string {
	dir := filepath.Dir(path)
	filename := filepath.Base(path) + ".rels"
	return filepath.Join(dir, "_rels", filename)
}

// TargetPath returns the relative path from source part file to target part file.
func TargetPath(src, dst Part) string {
	path, err := filepath.Rel(filepath.Dir(src.Path()), dst.Path())
	if err != nil {
		log.Fatal(err)
	}
	return path
}
