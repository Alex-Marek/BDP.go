package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetDisconnect(reader *bitreader.Reader) {
	fmt.Println("Parse Net Disconnect")
	text := reader.TryReadString()
	fmt.Printf("\t%s %s\n", "Text:", text)
}
