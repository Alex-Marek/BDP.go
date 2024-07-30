package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetSignonState(reader *bitreader.Reader) {
	fmt.Println("Parse Net Sign On State")
	SignonState := reader.TryReadUInt8()
	SpawnCount := reader.TryReadUInt32()
	NumServerPlayers := reader.TryReadUInt32()
	IdsLength := reader.TryReadUInt32()
	reader.SkipBytes(uint64(IdsLength))
	MapNameLength := reader.TryReadUInt32()
	MapName := reader.TryReadStringLength(uint64(MapNameLength))

	fmt.Printf("\t%s %d\n", " Sign On State:", SignonState)
	fmt.Printf("\t%s %d\n", " Spawn Count:", SpawnCount)
	fmt.Printf("\t%s %d\n", " Number of Server Players", NumServerPlayers)
	fmt.Printf("\t%s %s\n", " Map Name:", MapName)

}
