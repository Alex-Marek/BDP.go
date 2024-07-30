package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type conVar struct {
	Name  string
	Value string
}

func ParseNetSetConvar(reader *bitreader.Reader) {
	fmt.Println("Parse Net Set Convar")
	convars := []conVar{}
	length := reader.TryReadUInt8()
	for count := 0; count < int(length); count++ {
		convar := conVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		}
		convars = append(convars, convar)
	}
}
