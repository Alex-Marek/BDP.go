package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcServerInfo(reader *bitreader.Reader) {
	fmt.Println("Parsing Svc Server Info")
	Protocol := reader.TryReadUInt16()
	ServerCount := reader.TryReadUInt32()
	IsHltv := reader.TryReadBool()
	IsDedicated := reader.TryReadBool()
	ClientCrc := reader.TryReadUInt32()
	StringTableCrc := reader.TryReadUInt32()
	MaxClasses := reader.TryReadUInt16()
	MapCrc := reader.TryReadUInt32()
	PlayerSlot := reader.TryReadUInt8()
	MaxClients := reader.TryReadUInt8()
	TickInterval := reader.TryReadFloat32()
	OperatingSystem := reader.TryReadStringLength(1)
	GameDir := reader.TryReadString()
	MapName := reader.TryReadString()
	SkyName := reader.TryReadString()
	HostName := reader.TryReadString()

	fmt.Printf("\t%s %d\n", "Protocol:", Protocol)
	fmt.Printf("\t%s %d\n", "Server Count:", ServerCount)
	fmt.Printf("\t%s %t\n", "Is Hltv:", IsHltv)
	fmt.Printf("\t%s %t\n", "Is Dedicated:", IsDedicated)
	fmt.Printf("\t%s %d\n", "Client Crc:", ClientCrc)
	fmt.Printf("\t%s %d\n", "String Table Crc:", StringTableCrc)
	fmt.Printf("\t%s %d\n", "Max Server Classes:", MaxClasses)
	fmt.Printf("\t%s %d\n", "Map Crc:", MapCrc)
	fmt.Printf("\t%s %d\n", "Player Slot:", PlayerSlot)
	fmt.Printf("\t%s %d\n", "Max Clients:", MaxClients)
	fmt.Printf("\t%s %g\n", "Tick Interval:", TickInterval)
	fmt.Printf("\t%s %s\n", "Operating System:", OperatingSystem)
	fmt.Printf("\t%s %s\n", "Game Dir:", GameDir)
	fmt.Printf("\t%s %s\n", "Map Name:", MapName)
	fmt.Printf("\t%s %s\n", "Sky Name:", SkyName)
	fmt.Printf("\t%s %s\n", "Host Name:", HostName)
}
