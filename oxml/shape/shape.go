package shape

import (
	"encoding/xml"
)

// Shape is a kind of Drawing Charts such as a rectangle, circle, line and so on.
type Shape interface {
	MarshalXML(*xml.Encoder, xml.StartElement) error
}
