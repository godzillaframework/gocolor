package gocolor

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
