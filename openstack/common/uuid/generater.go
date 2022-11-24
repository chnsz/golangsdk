package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// The size of a UUID in bytes.
const size = 16

type UUID [size]byte

// Random is a method to generate a random UUID.
func Random() (UUID, error) {
	var u UUID
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return u, err
	}
	return u, nil
}

// String method returns a canonical string representation of UUID, the format is: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}
