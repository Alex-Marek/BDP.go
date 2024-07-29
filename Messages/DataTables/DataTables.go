package DataTables

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadDataTables(reader *bitreader.Reader) {
	fmt.Println("ReadDataTables Skips Most Content")
	size := reader.TryReadUInt32()
	reader.SkipBytes(uint64(size))
	fmt.Printf("%s %d\n", "Size Of DataTable:", size)
}
