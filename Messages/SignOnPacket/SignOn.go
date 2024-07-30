package SignOnPacket

import (
	"bdpv0/Messages/SignOnPacket/NetMessages"
	"bdpv0/Messages/SignOnPacket/SvcMessages"
	"fmt"

	"github.com/pektezol/bitreader"
)

const MSSC = 2
const parseCmdInfo = false

func ReadSignOnPacket(reader *bitreader.Reader) {
	fmt.Println("SignOn Or Packet")
	if parseCmdInfo == true {
		for count := 0; count < MSSC; count++ {
			ParseCmdInfo(reader)
		}
	} else {
		// 76 * MSSC
		reader.SkipBytes(152)
	}

	InSeq := reader.TryReadUInt32()
	OutSeq := reader.TryReadUInt32()
	size := reader.TryReadUInt32()

	fmt.Printf("\t%s %d\n", "In Sequence:", InSeq)
	fmt.Printf("\t%s %d\n", "Out Sequence:", OutSeq)
	fmt.Printf("\t%s %d\n", "Size:", size)

	NetSvcPacketReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(size)), true)
	for {
		NetSvcType, err := NetSvcPacketReader.ReadBits(6)
		if err != nil {
			break
		}
		switch NetSvcType {
		case 0:
			NetMessages.ParseNetNop()
		case 1:
			NetMessages.ParseNetDisconnect(NetSvcPacketReader)
		case 2:
			NetMessages.ParseNetFile(NetSvcPacketReader)
		case 3:
			NetMessages.ParseNetSplitScreenUser(NetSvcPacketReader)
		case 4:
			NetMessages.ParseNetTick(NetSvcPacketReader)
		case 5:
			NetMessages.ParseNetStringCmd(NetSvcPacketReader)
		case 6:
			NetMessages.ParseNetSetConvar(NetSvcPacketReader)
		case 7:
			NetMessages.ParseNetSignonState(NetSvcPacketReader)
		case 8:
			SvcMessages.ParseSvcServerInfo(NetSvcPacketReader)
		case 9:
			SvcMessages.ParseSvcSendTable(NetSvcPacketReader)
		case 10:
			SvcMessages.ParseSvcClassInfo(NetSvcPacketReader)
		case 11:
			SvcMessages.ParseSvcSetPause(NetSvcPacketReader)
		case 12:
			SvcMessages.ParseSvcCreateStringTable(NetSvcPacketReader)
		case 13:
			SvcMessages.ParseSvcUpdateStringTable(NetSvcPacketReader)
		case 14:
			SvcMessages.ParseSvcVoiceInit(NetSvcPacketReader)
		case 15:
			SvcMessages.ParseSvcVoiceData(NetSvcPacketReader)
		case 16:
			SvcMessages.ParseSvcPrint(NetSvcPacketReader)
		case 17:
			SvcMessages.ParseSvcSounds(NetSvcPacketReader)
		case 18:
			SvcMessages.ParseSvcSetView(NetSvcPacketReader)
		case 19:
			SvcMessages.ParseSvcFixAngle(NetSvcPacketReader)
		case 20:
			SvcMessages.ParseSvcCrosshairAngle(NetSvcPacketReader)
		case 21:
			SvcMessages.ParseBspDecal(NetSvcPacketReader)
		case 22:
			SvcMessages.ParseSvcSplitScreen(NetSvcPacketReader)
		case 23:
			// Will Probably Have Issues
			SvcMessages.ParseSvcUserMessage(NetSvcPacketReader)
		case 24:
			// Will Probably Have Issues
			SvcMessages.ParseEntityMessage(NetSvcPacketReader)
		case 25:
			// Will Probably Have Issues
			SvcMessages.ParseSvcGameEvent(NetSvcPacketReader)
		case 26:
			// Will Probably Have Issues
			SvcMessages.ParseSvcPacketEntities(NetSvcPacketReader)
		case 27:
			// Will Probably Have Issues
			SvcMessages.ParseSvcTempEntities(NetSvcPacketReader)
		case 28:
			SvcMessages.ParseSvcPrefetch(NetSvcPacketReader)
		case 29:
			SvcMessages.ParseSvcMenu(NetSvcPacketReader)
		case 30:
			SvcMessages.ParseGameEventList(NetSvcPacketReader)
		case 31:
			SvcMessages.ParseSvcGetCvarValue(NetSvcPacketReader)
		case 32:
			SvcMessages.ParseSvcCmdKeyValues(NetSvcPacketReader)
		case 33:
			SvcMessages.ParseSvcPaintmapData(NetSvcPacketReader)
		default:
			panic("Invalid NetSvcType")
		}
	}

}

func ParseCmdInfo(reader *bitreader.Reader) {
	Flags := reader.TryReadUInt32()
	ViewOrigin := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	ViewAngles := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	LocalViewAngles := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	ViewOrigin2 := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	ViewAngles2 := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}
	LocalViewAngles2 := []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()}

	fmt.Println("PacketInfo - Cmd Info: ")
	fmt.Printf("\t%s%d\n", " Flags: ", Flags)
	fmt.Printf("\t%s%g %g %g\n", " ViewOrigin: ", ViewOrigin[0], ViewOrigin[1], ViewOrigin[2])
	fmt.Printf("\t%s%g %g %g\n", " ViewAngles: ", ViewAngles[0], ViewAngles[1], ViewAngles[2])
	fmt.Printf("\t%s%g %g %g\n", " LocalViewAngles: ", LocalViewAngles[0], LocalViewAngles[1], LocalViewAngles[2])
	fmt.Printf("\t%s%g %g %g\n", " ViewOrigin2: ", ViewOrigin2[0], ViewOrigin2[1], ViewOrigin2[2])
	fmt.Printf("\t%s%g %g %g\n", " ViewAngles2: ", ViewAngles2[0], ViewAngles2[1], ViewAngles2[2])
	fmt.Printf("\t%s%g %g %g\n", " LocalViewAngles2: ", LocalViewAngles2[0], LocalViewAngles2[1], LocalViewAngles2[2])
}
