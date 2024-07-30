package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetFile(reader *bitreader.Reader) {
	fmt.Println("Parse Net File")
	reader.TryReadUInt32() // TransferId
	reader.TryReadString() // FileName
	reader.TryReadBool()   // File Requested (Bool)
}
