package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcTempEntities(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Temp Entities")
	NumEntries := reader.TryReadUInt8()
	Length := uint32(reader.TryReadBits(17))
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Number Of Entries:", NumEntries)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %v\n", "Data:", Data)
}
