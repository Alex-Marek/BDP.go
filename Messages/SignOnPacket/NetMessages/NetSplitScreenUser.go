package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetSplitScreenUser(reader *bitreader.Reader) {
	fmt.Println("Parse Net Split Screen User")
	reader.TryReadBool()
}
