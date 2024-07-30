package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcSendTable(reader *bitreader.Reader) {
	fmt.Println("Parsing Svc Send Table")
	NeedsDecoder := reader.TryReadBool()
	Length := reader.TryReadUInt8()
	Props := uint32(reader.TryReadBits(uint64(Length)))

	fmt.Printf("\t%s %t\n", "Needs Decoder:", NeedsDecoder)
	fmt.Printf("\t%s %d\n", "Props:", Props)
}
