package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseEntityMessage(reader *bitreader.Reader) {
	fmt.Println("Parse Entity Message")
	EntityIndex := uint16(reader.TryReadBits(11))
	ClassId := uint16(reader.TryReadBits(9))
	Length := uint16(reader.TryReadBits(11))
	Data := reader.TryReadBitsToSlice(uint64(Length))

	fmt.Printf("\t%s %d\n", "Entity Index:", EntityIndex)
	fmt.Printf("\t%s %d\n", "Class ID:", ClassId)
	fmt.Printf("\t%s %d\n", "Length:", Length)
	fmt.Printf("\t%s %x\n", "Data:", Data)
}
