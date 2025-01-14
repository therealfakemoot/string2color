package string2color

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"

	"golang.org/x/exp/constraints"
)

type RGB struct {
	R, G, B uint16
}

func (rgb RGB) String() string {
	return fmt.Sprintf("{R:%x,G:%x,B:%x}", rgb.R, rgb.G, rgb.B)
}

type RGBA struct {
	RGB
	A uint16
}

func (rgba RGBA) String() string {
	return fmt.Sprintf("{R:%x,G:%x,B:%x,A:%x}", rgba.R, rgba.G, rgba.B, rgba.A)
}

func ToRGB(s string) RGB {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	var rgb RGB
	r := binary.LittleEndian.Uint16(sum[:7])
	log.Printf("red value: %d\n", r)
	g := binary.LittleEndian.Uint16(sum[7:14])
	log.Printf("green value: %d\n", g)
	b := binary.LittleEndian.Uint16(sum[13:])
	log.Printf("blue value: %d\n", b)
	rgb.R = InterpolateUint16(r, 0, uint16(math.Pow(2, 7)), 0, 16384)
	rgb.G = InterpolateUint16(g, 0, uint16(math.Pow(2, 7)), 0, 16384)
	rgb.B = InterpolateUint16(b, 0, uint16(math.Pow(2, 7)), 0, 16384)

	return rgb
}

func ToRGBA(b []byte) RGBA {
	var rgba RGBA

	return rgba
}

func Interpolate[T constraints.Unsigned](f, inputMin, inputMax, outputMin, outputMax T) T {
	return (f-(inputMin))*(outputMax-outputMin)/(inputMax-inputMin) + outputMin
}

var InterpolateUint16 = Interpolate[uint16]

/*
func Interpolate(f, inputMin, inputMax, outputMin, outputMax uint16) uint16 {
	return (f-(inputMin))*(outputMax-outputMin)/(inputMax-inputMin) + outputMin
}
*/
