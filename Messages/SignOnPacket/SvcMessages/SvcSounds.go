package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcSounds(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Sounds")
	ReliableSound := reader.TryReadBool()
	var NumSounds uint8
	var Length uint16
	if ReliableSound {
		NumSounds = 1
		Length = uint16(reader.TryReadUInt8())
	} else {
		NumSounds = reader.TryReadUInt8()
		Length = uint16(reader.TryReadUInt16())
	}
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %t\n", "Reliable Sound:", ReliableSound)
	fmt.Printf("\t%s %d\n", "Number Of Sounds:", NumSounds)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %v\n", "Data:", Data)
}
