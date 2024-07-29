package UserCmd

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadUserCmd(reader *bitreader.Reader) {
	fmt.Println("UserCmd")
	cmd := reader.TryReadStringLength(4)
	size := reader.TryReadUInt32()
	reader.SkipBytes(uint64(size))

	fmt.Printf("%s %s\n", " Command", cmd)
	fmt.Printf("%s %d\n", " Size:", size)
}
