package string2color

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
)

func Convert(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	sum := h.Sum(nil)
	fmt.Println(binary.LittleEndian.Uint32(sum[:8]))

	return ""
}
