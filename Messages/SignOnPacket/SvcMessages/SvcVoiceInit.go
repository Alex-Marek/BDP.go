package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcVoiceInit(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Voice Init")
	Codec := reader.TryReadString()
	Quality := reader.TryReadUInt8()
	var SampleRate int32
	if Quality == 0b11111111 {
		SampleRate = reader.TryReadSInt32()
	} else {
		if Codec == "vaudio_celt" {
			SampleRate = 22050
		} else {
			SampleRate = 11025
		}
	}
	fmt.Printf("\t%s %s\n", "Codec: ", Codec)
	fmt.Printf("\t%s %d\n", "Quality: ", Quality)
	fmt.Printf("\t%s %d\n", "Sample Rate: ", SampleRate)
}
