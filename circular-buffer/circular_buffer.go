package circular

import (
	"errors"
)

var errBufferFull = errors.New("buffer is full")
var errBufferEmpty = errors.New("buffer is empty")

type Buffer struct {
	buf  []byte
	size int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{size: size}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.buf) == 0 {
		b.Reset()
		return 0, errBufferEmpty
	}
	c := b.buf[0]
	b.buf = b.buf[1:]
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buf) < b.size {
		b.buf = append(b.buf, c)
		return nil
	}
	return errBufferFull
}

func (b *Buffer) Overwrite(c byte) {
	if err := b.WriteByte(c); err != nil {
		b.buf = append(b.buf[1:], c)
	}
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}
