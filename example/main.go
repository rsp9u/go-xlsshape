package main

import (
	"github.com/rsp9u/go-xlsshape/oxml"
	"github.com/rsp9u/go-xlsshape/oxml/shape"
)

func main() {
	ss := oxml.NewSpreadsheet()

	rect := shape.NewRectangle()
	rect.SetLeftTop(20, 20)
	rect.SetSize(300, 150)
	rect.SetText("A shape with text", "en-US")
	ss.AddShape(rect)

	rect2 := shape.NewRectangle()
	rect2.SetLeftTop(320, 20)
	rect2.SetSize(300, 150)
	rect2.SetNoFill(true)
	rect2.SetNoLine(true)
	rect2.SetText("A shape with text\nWithout fill and line", "en-US")
	ss.AddShape(rect2)

	line1 := shape.NewLine()
	line1.SetStartPos(20, 300)
	line1.SetEndPos(300, 300)
	line1.SetTailType("arrow")
	ss.AddShape(line1)

	line2 := shape.NewLine()
	line2.SetStartPos(20, 340)
	line2.SetEndPos(300, 340)
	line2.SetDashType("dash")
	line2.SetHeadType("triangle")
	line2.SetTailType("stealth")
	ss.AddShape(line2)

	line3 := shape.NewLine()
	line3.SetStartPos(20, 380)
	line3.SetEndPos(300, 380)
	line3.SetDashType("sysDot")
	ss.AddShape(line3)

	ss.Dump("example1.xlsx")
}
