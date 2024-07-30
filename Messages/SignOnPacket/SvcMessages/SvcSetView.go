package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcSetView(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Set View")
	EntityIndex := reader.TryReadBits(11)
	fmt.Printf("\t%s %d\n", "Entity Index:", uint64(EntityIndex))
}
