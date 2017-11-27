// Package paasio provides functions about ReadCounter and WriteCounter
package paasio

import (
	"io"
	"sync"
)

// NewReadCounter receives an io.Reader and returns a ReadCounter
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

type readCounter struct {
	reader    io.Reader
	bytesRead int64
	nops      int
	sync.Mutex
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.reader.Read(p)

	rc.Lock()
	defer rc.Unlock()

	rc.bytesRead += int64(n)
	rc.nops++
	return n, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.bytesRead, rc.nops
}

// NewWriteCounter receives an io.Writer and returns a WriteCounter
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

type writeCounter struct {
	writer       io.Writer
	bytesWritten int64
	nops         int
	sync.Mutex
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.writer.Write(p)

	wc.Lock()
	defer wc.Unlock()

	wc.bytesWritten += int64(n)
	wc.nops++
	return n, err
}

func (rc *writeCounter) WriteCount() (n int64, nops int) {
	return rc.bytesWritten, rc.nops
}

// NewReadWriteCounter receives an io.ReadWriter and returns a ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(rw), NewWriteCounter(rw)}
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}
