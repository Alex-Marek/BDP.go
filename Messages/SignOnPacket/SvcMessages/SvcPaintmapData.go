package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcPaintmapData(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Paint Map Data")
	Length := reader.TryReadUInt32()
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %x\n", "Data:", Data)
}
