package main

import "bytes"

// ByteWriter is a custom writer that writes to a byte buffer.
type ByteWriter struct {
	buffer *bytes.Buffer
}

// Write implements the io.Writer interface.
func (b *ByteWriter) Write(p []byte) (n int, err error) {
	return b.buffer.Write(p)
}
