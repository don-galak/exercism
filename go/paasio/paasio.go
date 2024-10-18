package paasio

import (
	"io"
	"sync"
)

type counter struct {
	bytes uint64
	ops   uint32
	sync.Mutex
}

func (c *counter) increment(n int) {
	c.Lock()
	defer c.Unlock()
	c.bytes += uint64(n)
	c.ops++
}

func (c *counter) count() (int64, int) {
	c.Lock()
	defer c.Unlock()
	return int64(c.bytes), int(c.ops)
}

type readCounter struct {
	reader io.Reader
	counter
}

type writeCounter struct {
	writer io.Writer
	counter
}

type readWriteCounter struct {
	WriteCounter
	ReadCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewWriteCounter(readwriter), NewReadCounter(readwriter)}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.increment(n)
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.increment(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.count()
}
