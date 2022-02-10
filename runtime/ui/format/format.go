package format

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lunixbochs/vtclean"
	"strings"
)

const (
	selectedLeftBracketStr  = "║"
	selectedRightBracketStr = "╠"
	selectedFillStr         = "═"

	leftBracketStr  = "│"
	rightBracketStr = "├"
	fillStr         = "─"

	selectStr = " "
)

var (
	Header                func(...interface{}) string
	ColumnHeader          func(...interface{}) string
	Selected              func(...interface{}) string
	SelectedHeader        func(...interface{}) string
	StatusSelected        func(...interface{}) string
	StatusNormal          func(...interface{}) string
	StatusControlSelected func(...interface{}) string
	StatusControlNormal   func(...interface{}) string
	CompareTop            func(...interface{}) string
	CompareBottom         func(...interface{}) string
)

func init() {
	Selected = color.New(color.ReverseVideo).SprintFunc()
	SelectedHeader = color.New(color.FgCyan).SprintFunc()
	Header = color.New(color.FgWhite).SprintFunc()
	ColumnHeader = color.New(color.FgBlue).SprintFunc()
	StatusSelected = color.New(color.BgBlue, color.FgBlack).SprintFunc()
	StatusNormal = color.New(color.ReverseVideo).SprintFunc()
	StatusControlSelected = color.New(color.BgBlue, color.FgBlack, color.Bold).SprintFunc()
	StatusControlNormal = color.New(color.ReverseVideo, color.Bold).SprintFunc()
	CompareTop = color.New(color.BgBlue).SprintFunc()
	CompareBottom = color.New(color.BgGreen).SprintFunc()
}

func RenderNoHeader(width int, selected bool) string {
	if selected {
		return strings.Repeat(selectedFillStr, width)
	}
	return strings.Repeat(fillStr, width)
}

func RenderHeader(title string, width int, selected bool) string {
	if selected {
		body := fmt.Sprintf("%s%s ", selectStr, title)
		bodyLen := len(vtclean.Clean(body, false))
		repeatCount := width - bodyLen - 2
		if repeatCount < 0 {
			repeatCount = 0
		}
		line := fmt.Sprintf("%s%s%s%s\n", selectedLeftBracketStr, body, selectedRightBracketStr, strings.Repeat(selectedFillStr, repeatCount))
		return SelectedHeader(line)
	}
	body := Header(fmt.Sprintf(" %s ", title))
	bodyLen := len(vtclean.Clean(body, false))
	repeatCount := width - bodyLen - 2
	if repeatCount < 0 {
		repeatCount = 0
	}
	return fmt.Sprintf("%s%s%s%s\n", leftBracketStr, body, rightBracketStr, strings.Repeat(fillStr, repeatCount))
}

func RenderHelpKey(control, title string, selected bool) string {
	if selected {
		return StatusSelected("▏") + StatusControlSelected(control) + StatusSelected(" "+title+" ")
	} else {
		return StatusNormal("▏") + StatusControlNormal(control) + StatusNormal(" "+title+" ")
	}
}
