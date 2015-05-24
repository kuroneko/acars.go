package acars

import (
	"unicode"
)

// Applies a simplfieid Tcl String Split
//
// Hoppie's ACARS services which all heavily utilise that style.
//
// This only splits a single level.  Call against results to sub-split.
func CurlySplit(strin string) []string {
	rs := make([]string, 0)
	curlyDepth := 0

	var curTok []rune = []rune{}

	for _, curRune := range strin {
		if (curlyDepth == 0) {
			if unicode.IsSpace(curRune) {
				// end the existing token
				if len(curTok) > 0 {
					rs = append(rs, string(curTok))
					curTok = []rune{}
				}
			} else {
				switch (curRune) {
				case '{':
					curlyDepth++
				default:
					curTok = append(curTok, curRune)
				}
			}
		} else {
			// depth > 0.  track curlies only.
			if (curRune == '{') {
				curlyDepth++;
			}
			if (curRune == '}') {
				curlyDepth--;
			}
			if (curlyDepth > 0) {
				curTok = append(curTok, curRune)
			}
		}
	}
	if len(curTok) > 0 {
		rs = append(rs, string(curTok))
		curTok = []rune{}
	}
	return rs
}