package circular

import (
	"errors"
	"io"
)

const opRead readOp = -1
const opInvalid readOp = 0

var ErrTooLarge = errors.New("bytes.Buffer: too large")

type readOp int8
type Buffer struct {
	buf      []byte
	off      int
	lastRead readOp
	size     int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{size: size}
}

func (b *Buffer) empty() bool { return len(b.buf) <= b.off }

func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
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
	if len(b.buf) < b.size {
		println(string(c))
		b.buf = append(b.buf, c)
		return nil
	}
	return ErrTooLarge
}

func (b *Buffer) Overwrite(c byte) {
	if err := b.WriteByte(c); err != nil {
		b.buf = append(b.buf[1:], c)
	}
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}
