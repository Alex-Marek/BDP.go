package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcGetCvarValue(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Get Cvar Value:")
	Cookie := reader.TryReadUInt32()
	CvarName := reader.TryReadString()

	fmt.Printf("\t%s %d\n", "Cookie:", Cookie)
	fmt.Printf("\t%s %s\n", "Cvar Name:", CvarName)
}
