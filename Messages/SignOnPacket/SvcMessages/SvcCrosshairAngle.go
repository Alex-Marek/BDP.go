package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcCrosshairAngle(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Crosshair Angle")
	XAng := float32(reader.TryReadBits(16))
	YAng := float32(reader.TryReadBits(16))
	ZAng := float32(reader.TryReadBits(16))
	fmt.Printf("\t%s %g %g %g\n", "X Y Z:", XAng, YAng, ZAng)
}
