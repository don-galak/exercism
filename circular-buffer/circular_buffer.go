package circular

import (
	"errors"
	"io"
)

const opRead readOp = -1
const opInvalid readOp = 0

var ErrTooLarge = errors.New("bytes.Buffer: too large")

type readOp int8

// Define the Buffer type here.
type Buffer struct {
	buf      []byte
	off      int
	lastRead readOp
}

func NewBuffer(size int) *Buffer {
	return &Buffer{}
}

func (b *Buffer) empty() bool { return len(b.buf) <= b.off }

func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
		// Buffer is empty, reset to recover space.
		b.Reset()
		return 0, io.EOF
	}
	c := b.buf[b.off]
	b.off++
	b.lastRead = opRead
	return c, nil
}

func (b *Buffer) Len() int { return len(b.buf) - b.off }

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buf) <= cap(b.buf) {
		b.buf = append(b.buf, c)
		return nil
	}
	return ErrTooLarge
}

func (b *Buffer) Overwrite(c byte) {
	b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}
