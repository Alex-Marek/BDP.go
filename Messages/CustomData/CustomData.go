package CustomData

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadCustomData(reader *bitreader.Reader) {
	fmt.Println("CustomData")
	Unknown := reader.TryReadBytes(4)
	size := reader.TryReadUInt32()
	reader.SkipBytes(uint64(size))
	fmt.Printf("%s %d\n", " Unknown:", Unknown)
	fmt.Printf("%s %d\n", " Size:", size)
}
