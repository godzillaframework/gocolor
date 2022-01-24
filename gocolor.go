package gocolor

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

/**
sequence color
*/
const (
	startSeq = "\033["
	endSeq   = "\033[0m"
)

type printer struct {
	text string
}

var color map[string]string

var Out io.Writer

var (
	colorGroupRE *regexp.Regexp = regexp.MustCompile(`(\{\w*\}[^{}]+)`)
	colorPartRE  *regexp.Regexp = regexp.MustCompile(`{(\w*)}`)
	textPartRE   *regexp.Regexp = regexp.MustCompile(`^{\w*}(.*)`)
)

func (p *printer) In(color string) {
	p.text = "{" + color + "}" + p.text
	p.inFormat()
}

func (p *printer) inFormat() {
	matches := colorGroupRE.FindAllStringSubmatch(p.text, -1)

	for _, value := range matches {
		color := colorPartRE.FindStringSubmatch(value[0])[1]
		colorcode := getColor(color)

		text := textPartRE.FindStringSubmatchIndex(value[0])[1]
		clifmt := startSeq + colorcode + text + endSeq
		p.text = strings.Replace(p.text, value[0], clifmt, -1)
	}

	fmt.Fprintf(Out, p.text)

}

func getColor(color string) string {
	var colorcode string

	if value, ok := colors[color]; ok {
		colorcode = value
	} else {
		colorcode = colors["default"]
	}
}
