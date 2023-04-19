package paasio

import (
	"io"
	"sync"
)

type counter struct {
	bytes uint64
	ops   uint32
	mutex *sync.Mutex
}

func (c *counter) addBytes(n int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.bytes += uint64(n)
	c.ops++
}

func (c *counter) count() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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

func newCounter() counter {
	return counter{mutex: new(sync.Mutex)}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer, newCounter()}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader, newCounter()}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewWriteCounter(readwriter), NewReadCounter(readwriter)}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.addBytes(n)
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.addBytes(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.count()
}
