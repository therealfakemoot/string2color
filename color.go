package string2color

import (
	"crypto/sha1"
	"encoding/binary"
	"io"
)

func Convert(s string) float64 {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	i := binary.LittleEndian.Uint32(sum[:8])
	j := float64(i) / float64(65536)
	k := j * 360

	return k
}
