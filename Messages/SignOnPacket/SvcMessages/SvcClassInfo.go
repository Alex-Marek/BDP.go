package SvcMessages

import (
	"fmt"
	"math"

	"github.com/pektezol/bitreader"
)

func ParseSvcClassInfo(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Class Info")
	Length := reader.TryReadUInt16()
	CreateOnClient := reader.TryReadBool()
	if !CreateOnClient {
		for count := 0; count < int(Length); count++ {
			ClassId := int16(reader.TryReadBits(uint64(math.Log2(float64(Length)) + 1)))
			ClassName := reader.TryReadString()
			DataTableName := reader.TryReadString()
			fmt.Printf("\t%s %d\n", "Class ID:", ClassId)
			fmt.Printf("\t%s %s\n", "Class Name:", ClassName)
			fmt.Printf("\t%s %s\n", "Data Table Name:", DataTableName)
		}
	}

}
