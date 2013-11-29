package testutil

import "testing"

// Silly little wrapper to provide a couple interfaces:
//  io.Writer
type IOWrapper struct {
	*testing.T
}

func (t IOWrapper) Write(buf []byte) (int, error) {
	if len(buf) > 0 && buf[len(buf)-1] == '\n' {
		buf = buf[:len(buf)-1]
	}
	t.Log(string(buf))
	return len(buf), nil
}
