// Package circular implements a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
//   type Buffer
//   func NewBuffer(size int) *Buffer
//   func (*Buffer) ReadByte() (byte, error)
//   func (*Buffer) WriteByte(c byte) error
//   func (*Buffer) Overwrite(c byte)
//   func (*Buffer) Reset() // put buffer in an empty state
//
// We chose the above API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.
package circular

import "errors"

const testVersion = 4

type Buffer struct {
	data              []byte
	size, head, count int
}

func NewBuffer(size int) *Buffer {
	if size <= 0 {
		return nil
	}
	var buffer Buffer
	buffer.data = make([]byte, size)
	buffer.size = size
	buffer.head = 0
	buffer.count = 0
	return &buffer
}

func (buf *Buffer) ReadByte() (ret byte, err error) {
	if buf.count > 0 {
		ret = buf.data[buf.head]
		buf.head++
		if buf.head >= buf.size {
			buf.head -= buf.size
		}
		buf.count--
	} else {
		err = errors.New("Empty buffer")
	}
	return
}

func (buf *Buffer) WriteByte(c byte) error {
	if buf.count == buf.size {
		return errors.New("No enough room for new data")
	}
	insertPos := buf.head + buf.count
	if insertPos >= buf.size {
		insertPos -= buf.size
	}
	buf.data[insertPos] = c
	buf.count++
	return nil
}

func (buf *Buffer) Overwrite(c byte) {
	if buf.count == buf.size {
		buf.data[buf.head] = c
		buf.head++
		if buf.head >= buf.size {
			buf.head -= buf.size
		}
		return
	}
	buf.WriteByte(c)
}

func (buf *Buffer) Reset() {
	buf.head = 0
	buf.count = 0
}
