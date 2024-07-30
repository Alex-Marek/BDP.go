package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcPrint(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Print")
	Message := reader.TryReadString()
	fmt.Printf("\t%s %s\n", "Message:", Message)
}
