package oxml

import (
	"bytes"
	"encoding/xml"
)

// Part is a part of OpenXML SpreadSheetPackage.
type Part interface {
	Path() string
	Content() string
}

// DefaultEncode returns an encoded xml string.
func DefaultEncode(object interface{}) (string, error) {
	buf := new(bytes.Buffer)
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	if err := enc.Encode(object); err != nil {
		return "", err
	}
	return xml.Header + buf.String(), nil
}
