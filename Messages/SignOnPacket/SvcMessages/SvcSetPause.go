package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcSetPause(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Set Pause")
	reader.TryReadBool()
}
