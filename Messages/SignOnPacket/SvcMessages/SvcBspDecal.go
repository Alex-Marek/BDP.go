package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type vectorCoord struct {
	Value float32
	Valid bool
}

func ParseBspDecal(reader *bitreader.Reader) {
	fmt.Println("Parse BSP Decal")
	_ = readVectorCoords(reader)
	DecalTextureIndex := int16(reader.TryReadBits(9))
	if reader.TryReadBool() {
		EntityIndex := uint16(reader.TryReadBits(11))
		ModelIndex := uint16(reader.TryReadBits(11))
		fmt.Printf("\t%s %d\n", "Entity Index:", EntityIndex)
		fmt.Printf("\t%s %d\n", "Model Index:", ModelIndex)
	}
	LowPriority := reader.TryReadBool()
	fmt.Printf("\t%s %d\n", "Decal Texture Index:", DecalTextureIndex)
	fmt.Printf("\t%s %t\n", "Low Priority", LowPriority)
}

// This function is straight lifted from Bisaxa I should rewrite it later so I'm not stealing lmao
func readVectorCoords(reader *bitreader.Reader) []vectorCoord {
	const COORD_INTEGER_BITS uint8 = 14
	const COORD_FRACTIONAL_BITS uint8 = 5
	const COORD_DENOMINATOR uint8 = 1 << COORD_FRACTIONAL_BITS
	const COORD_RESOLUTION float32 = 1.0 / float32(COORD_DENOMINATOR)
	readVectorCoord := func() float32 {
		value := float32(0)
		integer := reader.TryReadBits(1)
		fraction := reader.TryReadBits(1)
		if integer != 0 || fraction != 0 {
			sign := reader.TryReadBits(1)
			if integer != 0 {
				integer = reader.TryReadBits(uint64(COORD_INTEGER_BITS)) + 1
			}
			if fraction != 0 {
				fraction = reader.TryReadBits(uint64(COORD_FRACTIONAL_BITS))
			}
			value = float32(integer) + float32(fraction)*COORD_RESOLUTION
			if sign != 0 {
				value = -value
			}
		}
		return value
	}
	x := reader.TryReadBits(1)
	y := reader.TryReadBits(1)
	z := reader.TryReadBits(1)
	return []vectorCoord{{Value: readVectorCoord(), Valid: x != 0}, {Value: readVectorCoord(), Valid: y != 0}, {Value: readVectorCoord(), Valid: z != 0}}
}
