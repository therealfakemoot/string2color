package string2color

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"golang.org/x/exp/constraints"
)

type RGB struct {
	R, G, B uint64
}

func (rgb RGB) String() string {
	return fmt.Sprintf("{R:%x,G:%x,B:%x}", rgb.R, rgb.G, rgb.B)
}

type RGBA struct {
	RGB
	A uint64
}

func (rgba RGBA) String() string {
	return fmt.Sprintf("{R:%x,G:%x,B:%x,A:%x}", rgba.R, rgba.G, rgba.B, rgba.A)
}

func ToRGB(s string) RGB {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	var rgb RGB
	r := binary.LittleEndian.Uint64(sum[:9])
	g := binary.LittleEndian.Uint64(sum[7:16])
	b := binary.LittleEndian.Uint64(sum[11:])
	rgb.R = InterpolateUint64(r, 0, uint64(math.Pow(2, 56)), 0, 64384)
	rgb.G = InterpolateUint64(g, 0, uint64(math.Pow(2, 56)), 0, 64384)
	rgb.B = InterpolateUint64(b, 0, uint64(math.Pow(2, 56)), 0, 64384)

	return rgb
}

func ToRGBA(s string) RGBA {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	var rgba RGBA
	r := binary.LittleEndian.Uint64(sum[:9])
	g := binary.LittleEndian.Uint64(sum[5:14])
	b := binary.LittleEndian.Uint64(sum[9:18])
	a := binary.LittleEndian.Uint64(sum[12:])
	rgba.R = InterpolateUint64(r, 0, uint64(math.Pow(2, 56)), 0, 64384)
	rgba.G = InterpolateUint64(g, 0, uint64(math.Pow(2, 56)), 0, 64384)
	rgba.B = InterpolateUint64(b, 0, uint64(math.Pow(2, 56)), 0, 64384)
	rgba.A = InterpolateUint64(a, 0, uint64(math.Pow(2, 56)), 0, 64384)

	return rgba
}

func Interpolate[T constraints.Unsigned](f, inputMin, inputMax, outputMin, outputMax T) T {
	return (f-(inputMin))*(outputMax-outputMin)/(inputMax-inputMin) + outputMin
}

var (
	InterpolateUint16 = Interpolate[uint16]
	InterpolateUint32 = Interpolate[uint32]
	InterpolateUint64 = Interpolate[uint64]
)

/*
func Interpolate(f, inputMin, inputMax, outputMin, outputMax uint16) uint16 {
	return (f-(inputMin))*(outputMax-outputMin)/(inputMax-inputMin) + outputMin
}
*/
