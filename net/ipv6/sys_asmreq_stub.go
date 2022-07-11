// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows

package ipv6

import (
	"net"

	"github.com/hxx258456/ccgo/net/internal/socket"
)

func (so *sockOpt) setIPMreq(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	return errNotImplemented
}