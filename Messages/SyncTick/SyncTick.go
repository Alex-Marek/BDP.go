package SyncTick

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ReadSyncTick(reader *bitreader.Reader) {
	fmt.Println("Message: SyncTick")
}
