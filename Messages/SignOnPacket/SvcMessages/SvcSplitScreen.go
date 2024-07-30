package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcSplitScreen(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Split Screen")
	Type := reader.TryReadUInt8()
	Length := uint16(reader.TryReadBits(11))
	Data := reader.TryReadBits(uint64(Length))
	fmt.Printf("\t%s %d", "Type:", Type)
	fmt.Printf("\t%s %d", "Length:", Length)
	fmt.Printf("\t%s %v", "Data:", Data)
}
