package osc

import (
	"context"
	"errors"
	"net"
	"strings"
)

const (
	// bufSize is the size of read and write buffers.
	// SuperCollider synthdef messages can easily have as much as 64K of data.
	bufSize = 65536
)

// Common errors.
var (
	ErrNilDispatcher  = errors.New("nil dispatcher")
	ErrPrematureClose = errors.New("server cannot be closed before calling Listen")
)

// Conn defines the methods
type Conn interface {
	net.Conn

	Context() context.Context
	Serve(int, Dispatcher) error
	Send(Packet) error
	SendTo(net.Addr, Packet) error
}

var (
	invalidAddressRunes      = []rune{'#', ' '}
	invalidExactAddressRunes = append([]rune{'*', '?', '{', '}', ',', '[', ']'}, invalidAddressRunes...)
)

// ValidateAddress returns an error if addr contains
// characters that are disallowed by the OSC spec.
// If exactMatch is true, then the check is relaxed to allow pattern matching
// characters.
func ValidateAddress(addr string, exactMatch bool) error {
	var invalidRunes []rune
	if exactMatch {
		invalidRunes = invalidExactAddressRunes
	} else {
		// TODO: validate pattern correctness - e.g. balanced brackets, etc.
		invalidRunes = invalidAddressRunes
	}
	for _, chr := range invalidRunes {
		if strings.ContainsRune(addr, chr) {
			return ErrInvalidAddress
		}
	}
	return nil
}
