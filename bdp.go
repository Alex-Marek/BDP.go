package main

import (
	"bdpv0/Messages/ConsoleCmd"
	"bdpv0/Messages/CustomData"
	"bdpv0/Messages/DataTables"
	"bdpv0/Messages/Header"
	"bdpv0/Messages/SignOnPacket"
	"bdpv0/Messages/Stop"
	"bdpv0/Messages/StringTables"
	"bdpv0/Messages/SyncTick"
	"bdpv0/Messages/UserCmd"
	"fmt"

	"os"

	"github.com/pektezol/bitreader"
)

const littleendian = true

func main() {
	count := 1
	file, err := os.Open("StoptheBox_1595_76561198083196477_295601.dem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bitreader.NewReader(file, littleendian)

	Header.ReadHeader(reader)

	//Loop to read through messages
	for {
		PacketType := reader.TryReadUInt8()
		TickNumber := reader.TryReadSInt32()
		SlotNumber := reader.TryReadUInt8()
		fmt.Printf("%s %d\n \t%s %x\n \t%s %d\n \t%s %d\n",
			"Message Info for Message Number:",
			count,
			"Packet Type: ", PacketType,
			"TickNumber: ", TickNumber,
			"SlotNumber: ", SlotNumber)
		switch PacketType {
		case 1:
			SignOnPacket.ReadSignOnPacket(reader)
		case 2:
			SignOnPacket.ReadSignOnPacket(reader)
		case 3:
			SyncTick.ReadSyncTick(reader)
		case 4:
			ConsoleCmd.ReadConsoleCmd(reader)
		case 5:
			UserCmd.ReadUserCmd(reader)
		case 6:
			DataTables.ReadDataTables(reader)
		case 7:
			Stop.ReadStop(reader)
		case 8:
			CustomData.ReadCustomData(reader)
		case 9:
			StringTables.ReadStringTables(reader)
		default:
			panic("Invalid Packet Type")
		}
		count++
		fmt.Println()
	}
}
