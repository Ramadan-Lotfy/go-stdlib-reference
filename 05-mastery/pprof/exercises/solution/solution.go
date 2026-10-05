// Package solution holds the reference solutions for the runtime/pprof and net/http/pprof exercises.
package solution

import (
	"bytes"
	"runtime/pprof"
)

// ProfileCPU runs f while recording a CPU profile and returns the profile bytes
// (gzip-compressed protobuf, ready for `go tool pprof`). It returns an error if
// a CPU profile is already running.
func ProfileCPU(f func()) ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, err
	}
	defer pprof.StopCPUProfile()
	f()
	pprof.StopCPUProfile() // finish writing before reading buf
	return buf.Bytes(), nil
}

// HeapProfile returns the current heap profile in the binary format.
func HeapProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("heap").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
