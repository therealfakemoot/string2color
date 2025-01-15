package string2color

import (
	"crypto/sha1"
	"encoding/binary"
	"image/color"
	"io"
	"math"

	"golang.org/x/exp/constraints"
)

func ToRGB(s string) color.RGBA {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	var rgb color.RGBA
	r := binary.LittleEndian.Uint64(sum[:9])
	g := binary.LittleEndian.Uint64(sum[7:16])
	b := binary.LittleEndian.Uint64(sum[11:])
	rgb.R = uint8(InterpolateUint64(r, 0, uint64(math.Pow(2, 56)), 0, 64384))
	rgb.G = uint8(InterpolateUint64(g, 0, uint64(math.Pow(2, 56)), 0, 64384))
	rgb.B = uint8(InterpolateUint64(b, 0, uint64(math.Pow(2, 56)), 0, 64384))

	return rgb
}

func ToRGBA(s string) color.RGBA {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	var rgba color.RGBA
	r := binary.LittleEndian.Uint64(sum[:9])
	g := binary.LittleEndian.Uint64(sum[5:14])
	b := binary.LittleEndian.Uint64(sum[9:18])
	a := binary.LittleEndian.Uint64(sum[12:])
	rgba.R = uint8(InterpolateUint64(r, 0, uint64(math.Pow(2, 56)), 0, 64384))
	rgba.G = uint8(InterpolateUint64(g, 0, uint64(math.Pow(2, 56)), 0, 64384))
	rgba.B = uint8(InterpolateUint64(b, 0, uint64(math.Pow(2, 56)), 0, 64384))
	rgba.A = uint8(InterpolateUint64(a, 0, uint64(math.Pow(2, 56)), 0, 64384))

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
