package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcUpdateStringTable(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Update String Table")
	TableId := uint8(reader.TryReadBits(5))
	if reader.TryReadBool() {
		NumChangedEntries := reader.TryReadUInt16()
		fmt.Printf("\t%s %d\n", "Number of Entries Changed:", NumChangedEntries)
	}
	Length := int32(reader.TryReadBits(20))
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Table Id:", TableId)

	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %v\n", "Data:", Data)
}
