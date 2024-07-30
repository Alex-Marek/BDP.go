package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetStringCmd(reader *bitreader.Reader) {
	fmt.Println("Parse Net String Cmd")
	Command := reader.TryReadString()
	fmt.Printf("\t%s %s\n", "Command:", Command)
}
