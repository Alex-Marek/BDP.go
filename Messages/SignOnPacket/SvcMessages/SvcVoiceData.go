package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcVoiceData(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Voice Data")
	Client := reader.TryReadUInt8()
	Proximity := reader.TryReadUInt8()
	Length := reader.TryReadUInt16()
	// Audible Not parsed just lumped in with Data for now:(
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Client:", Client)
	fmt.Printf("\t%s %d\n", "Proximity:", Proximity)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %v\n", "Data:", Data)
}
