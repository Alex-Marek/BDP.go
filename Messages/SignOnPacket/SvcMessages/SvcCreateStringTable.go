package SvcMessages

import (
	"fmt"
	"math"

	"github.com/pektezol/bitreader"
)

func ParseSvcCreateStringTable(reader *bitreader.Reader) {
	fmt.Println("Parsing Svc Create String Table")

	Name := reader.TryReadString()
	MaxEntries := reader.TryReadSInt16()
	NumEntries := int8(reader.TryReadBits(uint64(math.Log2(float64(MaxEntries))) + 1))
	Length := int32(reader.TryReadBits(20))
	UserDataFixedSize := reader.TryReadBool()
	fmt.Printf("\t%s %s\n", "String Table Name:", Name)
	fmt.Printf("\t%s %d\n", "String Table Max Entries:", MaxEntries)
	fmt.Printf("\t%s %d\n", "String Table Number of Entries:", NumEntries)
	fmt.Printf("\t%s %t\n", "String Table User Data Fixed Size:", UserDataFixedSize)
	fmt.Printf("\t%s %d\n", "String Table Length:", Length)
	if UserDataFixedSize {
		UserDataSize := int16(reader.TryReadBits(12))
		UserDataSizeBits := int8(reader.TryReadBits(4))
		fmt.Printf("\t%s %d\n", "String Table User Data Size:", UserDataSize)
		fmt.Printf("\t%s %d\n", "String Table User Data Size Bit:", UserDataSizeBits)
	}
	Flags := int8(reader.TryReadBits(2))
	fmt.Printf("\t%s %d\n", "String Table Flags:", Flags)
	reader.SkipBits(uint64(Length)) // String Data
}
