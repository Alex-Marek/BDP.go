package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcGameEvent(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Game Event")
	Length := uint16(reader.TryReadBits(11))
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %x\n", "Data:", Data)
}
