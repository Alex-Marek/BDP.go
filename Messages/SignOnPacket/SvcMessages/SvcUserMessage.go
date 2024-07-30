package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcUserMessage(reader *bitreader.Reader) {
	fmt.Println("Parse Svc User Message")
	MsgType := int8(reader.TryReadBits(8))
	Length := int16(reader.TryReadBits(12))
	Data := reader.TryReadBitsToSlice(uint64(Length))
	fmt.Printf("\t%s %d\n", "Message Type:", MsgType)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %x\n", "Data:", Data)
}
