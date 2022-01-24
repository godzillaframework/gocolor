package gocolor

import (
	"fmt"
	"github.com/shiena/ansicolor"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
)

const (
	startSeq = "\033["
	endSeq   = "\033[0m"
)

type textPrint struct {
	text string
}

var colors map[string]string

var Out io.Writer

var (
	colorGroupRE *regexp.Regexp = regexp.MustCompile(`(\{\w*\}[^{}]+)`)
	colorPartRE  *regexp.Regexp = regexp.MustCompile(`{(\w*)}`)
	textPartRE   *regexp.Regexp = regexp.MustCompile(`^{\w*}(.*)`)
)

func (p *textPrint) In(color string) {
	p.text = "{" + color + "}" + p.text
	p.inFormat()
}

func (p *textPrint) inFormat() {
	matches := colorGroupRE.FindAllStringSubmatch(p.text, -1)

	for _, value := range matches {
		color := colorPartRE.FindStringSubmatch(value[0])[1]
		colorcode := getColor(color)

		text := textPartRE.FindStringSubmatch(value[0])[1]

		clifmt := startSeq + colorcode + text + endSeq
		p.text = strings.Replace(p.text, value[0], clifmt, -1)
	}

	fmt.Fprintln(Out, p.text)
}

func init() {
	colors = make(map[string]string)
	colors["black"] = "30m"
	colors["red"] = "31m"
	colors["green"] = "32m"
	colors["yellow"] = "33m"
	colors["blue"] = "34m"
	colors["magenta"] = "35m"
	colors["cyan"] = "36m"
	colors["white"] = "37m"
	colors["default"] = "39m"

	if runtime.GOOS == "windows" {
		Out = ansicolor.NewAnsiColorWriter(os.Stdout)
	} else {
		Out = os.Stdout
	}
}

func getColor(color string) string {
	var colorcode string

	if value, ok := colors[color]; ok {
		colorcode = value
	} else {
		colorcode = colors["default"]
	}

	return colorcode
}

func Print(color string) *textPrint {
	return &textPrint{color}
}
