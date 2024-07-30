package SvcMessages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseSvcFixAngle(reader *bitreader.Reader) {
	fmt.Println("Parse Svc Fix Angle")
	Relative := reader.TryReadBool()
	XAng := float32(reader.TryReadBits(16))
	YAng := float32(reader.TryReadBits(16))
	ZAng := float32(reader.TryReadBits(16))
	fmt.Printf("\t%s %t\n", "Relative:", Relative)
	fmt.Printf("\t%s %g %g %g\n", "X Y Z:", XAng, YAng, ZAng)
}
