package acars

// A CpdlcString is an Hoppie's ACARS encoded CPDLC message with it's native encoding
type CpdlcString string

// String returns the platform neutral encoding
func (str CpdlcString) String() string {
	rOut := []rune{}

	for _, r := range str {
		switch r {
		case '/':
			// Hoppie ACARS internal delimiter.  Not valid in cpdlcstrings. discard.
		case '|':
			rOut = append(rOut, '/')
		case '@':
			rOut = append(rOut, '\n')
		default:
			rOut = append(rOut, r)
		}
	}
	return string(rOut)
}

// NewCpdlcString encodes a platform neutral string into Hoppie's CPDLC format
func NewCpdlcString(strIn string) CpdlcString {
	rOut := []rune{}

	var consumeNewline = false
	for _, r := range strIn {
		switch r {
		case '@', '|':
			// unrepresentable characters
			rOut = append(rOut, '?')
			consumeNewline = false
		case '/':
			rOut = append(rOut, '|')
			consumeNewline = false
		case '\r', '\n':
			if !consumeNewline {
				rOut = append(rOut, '@')
				consumeNewline = false
			}
			consumeNewline = true
		default:
			rOut = append(rOut, r)
			consumeNewline = false
		}
	}
	return CpdlcString(rOut)
}
