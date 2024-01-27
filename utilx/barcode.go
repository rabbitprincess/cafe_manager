package utilx

import (
	"encoding/hex"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/codabar"
)

func Barcode(metadata string) (barcode.Barcode, error) {
	hex.EncodeToString([]byte(metadata))
	return codabar.Encode(metadata)
}
