package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcMenu(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Menu")
	MenuType := reader.TryReadUInt16()
	Length := reader.TryReadUInt32()
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d", "Menu Type:", MenuType)
	fmt.Printf("\t%s %d", "Length:", Length)
	fmt.Printf("\t%s %x", "Data:", Data)

}
