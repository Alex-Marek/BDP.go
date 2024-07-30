package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseGameEventList(reader *bitreader.Reader) {
	fmt.Println("Parse Game Event List")
	Events := uint16(reader.TryReadBits(9))
	Length := uint32(reader.TryReadBits(20))
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d", "Events:", Events)
	fmt.Printf("\t%s %d", "Length:", Length)
	fmt.Printf("\t%s %x", "Data:", Data)
}
