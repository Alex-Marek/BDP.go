package Header

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadHeader(reader *bitreader.Reader) {
	DemoFileStamp := reader.TryReadString()
	DemoProtocol := int32(reader.TryReadSInt32())
	NetworkProtocol := int32(reader.TryReadSInt32())
	ServerName := reader.TryReadStringLength(260)
	ClientName := reader.TryReadStringLength(260)
	MapName := reader.TryReadStringLength(260)
	GameDirectory := reader.TryReadStringLength(260)
	PlaybackTime := reader.TryReadFloat32()
	PlaybackTicks := int32(reader.TryReadSInt32())
	PlaybackFrames := int32(reader.TryReadSInt32())
	SignOnLength := int32(reader.TryReadSInt32())

	fmt.Printf("%s\n %s %s\n %s %d\n %s %d\n %s %s\n %s %s\n %s %s\n %s %s\n %s %g\n %s %d\n %s %d\n %s %d\n\n",
		"Demo Header Info: ",
		"Demo File Stamp:", DemoFileStamp,
		"Demo Protocol:", DemoProtocol,
		"Network Protocol:", NetworkProtocol,
		"Server Name:", ServerName,
		"Client Name:", ClientName,
		"Map Name:", MapName,
		"GameDir:", GameDirectory,
		"PlayBackTime:", PlaybackTime,
		"PlayBackTicks:", PlaybackTicks,
		"PlayBackFrames:", PlaybackFrames,
		"SignOnLength:", SignOnLength)
}
