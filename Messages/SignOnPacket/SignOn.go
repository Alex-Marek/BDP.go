package SignOnPacket

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

const MSSC = 2
const parseCmdInfo = false

func ReadSignOnPacket(reader *bitreader.Reader) {
	fmt.Println("SignOnPacket")
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
	reader.SkipBytes(uint64(size))

	fmt.Printf("%s %d\n", "In Sequence:", InSeq)
	fmt.Printf("%s %d\n", "Out Sequence:", OutSeq)
	fmt.Printf("%s %d\n", "Size:", size)
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
	fmt.Printf("%s%d\n", " Flags: ", Flags)
	fmt.Printf("%s%g %g %g\n", " ViewOrigin: ", ViewOrigin[0], ViewOrigin[1], ViewOrigin[2])
	fmt.Printf("%s%g %g %g\n", " ViewAngles: ", ViewAngles[0], ViewAngles[1], ViewAngles[2])
	fmt.Printf("%s%g %g %g\n", " LocalViewAngles: ", LocalViewAngles[0], LocalViewAngles[1], LocalViewAngles[2])
	fmt.Printf("%s%g %g %g\n", " ViewOrigin2: ", ViewOrigin2[0], ViewOrigin2[1], ViewOrigin2[2])
	fmt.Printf("%s%g %g %g\n", " ViewAngles2: ", ViewAngles2[0], ViewAngles2[1], ViewAngles2[2])
	fmt.Printf("%s%g %g %g\n", " LocalViewAngles2: ", LocalViewAngles2[0], LocalViewAngles2[1], LocalViewAngles2[2])
}
