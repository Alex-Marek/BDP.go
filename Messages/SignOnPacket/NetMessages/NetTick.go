package NetMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseNetTick(reader *bitreader.Reader) {
	fmt.Println("Parse Net Tick")
	Tick := reader.TryReadUInt32()                      // Tick
	HostFrameTime := reader.TryReadUInt16()             // HostFrameTime
	HostFrameTimeStdDeviation := reader.TryReadUInt16() // HostFrameTimeStdDeviation
	fmt.Printf("\t%s %d\n", "Engine Tick:", Tick)
	fmt.Printf("\t%s %f\n", "Engine Tick:", float32(HostFrameTime)/1e5)
	fmt.Printf("\t%s %f\n", "Engine Tick:", float32(HostFrameTimeStdDeviation)/1e5)
}
