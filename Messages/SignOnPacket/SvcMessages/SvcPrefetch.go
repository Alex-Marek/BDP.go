package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcPrefetch(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Prefetch")
	SoundIndex := uint16(reader.TryReadBits(13))
	fmt.Printf("\t%s %d\n", "Sound Index:", SoundIndex)
}
