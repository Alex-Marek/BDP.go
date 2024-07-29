package StringTables

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadStringTables(reader *bitreader.Reader) {
	fmt.Println("StringTables")
	size := reader.TryReadUInt32()
	reader.SkipBytes(uint64(size))
	fmt.Printf("%s %d\n", " Size:", size)
}
