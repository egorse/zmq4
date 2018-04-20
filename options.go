// Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zmq4

import (
	"time"

	"github.com/go-zeromq/zmq4/zmtp"
	"github.com/go-zeromq/zmq4/zmtp/security/null"
)

// Option configures some aspect of a ZeroMQ socket.
// (e.g. SocketIdentity, Security, ...)
type Option func(s *socket)

// WithID configures a ZeroMQ socket identity.
func WithID(id zmtp.SocketIdentity) func(s *socket) {
	return func(s *socket) {
		s.id = id
	}
}

// WithSecurity configures a ZeroMQ socket to use the given security mechanism.
// If the security mechanims is nil, the NULL mechanism is used.
func WithSecurity(sec zmtp.Security) func(s *socket) {
	return func(s *socket) {
		if sec == nil {
			sec = null.Security()
		}
		s.sec = sec
	}
}

// WithDialerRetry configures the time to wait before two failed attempts
// at dialing an endpoint.
func WithDialerRetry(retry time.Duration) func(s *socket) {
	return func(s *socket) {
		s.retry = retry
	}
}

// WithDialerTimeout sets the maximum amount of time a dial will wait
// for a connect to complete.
func WithDialerTimeout(timeout time.Duration) func(s *socket) {
	return func(s *socket) {
		s.dialer.Timeout = timeout
	}
}

/*
// TODO(sbinet)

func WithIOThreads(threads int)  func(s *socket) {
	return nil
}

func WithSendBufferSize(size int) func(s *socket) {
	return nil
}

func WithRecvBufferSize(size int) func(s *socket) {
	return nil
}
*/
