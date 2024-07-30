package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcCmdKeyValues(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Cmd Key Values")
	Length := reader.TryReadUInt32()
	Data := reader.TryReadBytesToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %v\n", "Data:", Data)
}
