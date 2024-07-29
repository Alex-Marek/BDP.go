package ConsoleCmd

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadConsoleCmd(reader *bitreader.Reader) {
	fmt.Println("Console Command")
	size := reader.TryReadUInt32()
	command := reader.TryReadStringLength(uint64(size))

	fmt.Printf("%s %s\n", " Command", command)
	fmt.Printf("%s %d\n", " Size:", size)
}
