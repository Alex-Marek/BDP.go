package Stop

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadStop(reader *bitreader.Reader) {
	fmt.Println("Stop")
	RemainingData := reader.TryReadBitsToSlice(uint64(reader.TryReadRemainingBits()))
	fmt.Printf("%s %v\n", " Remaining Data:", RemainingData)
}
