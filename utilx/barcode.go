package utilx

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/codabar"
)

func Barcode(metadata string) (barcode.Barcode, error) {
	return codabar.Encode(metadata)
}
