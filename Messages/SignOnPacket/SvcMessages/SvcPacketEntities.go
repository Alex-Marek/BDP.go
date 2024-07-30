package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcPacketEntities(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Packet Entities")
	MaxEntries := uint16(reader.TryReadBits(11))
	IsDelta := reader.TryReadBool()
	if IsDelta {
		DeltaFrom := reader.TryReadSInt32()
		fmt.Printf("\t%s %d\n", "Delta From:", DeltaFrom)
	}
	BaseLine := reader.TryReadBool()
	UpdatedEntries := uint16(reader.TryReadBits(11))
	Length := uint32(reader.TryReadBits(20))
	UpdateBaseLine := reader.TryReadBool()
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Max Entries:", MaxEntries)
	fmt.Printf("\t%s %t\n", "Is Delta:", IsDelta)
	fmt.Printf("\t%s %t\n", "Is Base Line:", BaseLine)
	fmt.Printf("\t%s %d\n", "Updated Entries:", UpdatedEntries)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %t\n", "Is Updated Base Line:", UpdateBaseLine)
	fmt.Printf("\t%s %x\n", "Data:", Data)
}
