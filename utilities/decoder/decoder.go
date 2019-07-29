package decoder

import (
	"golang.org/x/text/encoding/traditionalchinese"
)

/*
DecodeToBig5 deocdes string to big5
*/
func DecodeToBig5(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := traditionalchinese.Big5.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}
