package gocolor

import (
	"io"
	"regexp"
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
