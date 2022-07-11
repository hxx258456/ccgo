// Copyright 2022 s1ren@github.com/hxx258456.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

/*
gmtls是基于`golang/go`的`tls`包实现的国密改造版本。
对应版权声明: thrid_licenses/github.com/golang/go/LICENSE
*/

package gmtls

import (
	"errors"
	"syscall"
)

func init() {
	isConnRefused = func(err error) bool {
		return errors.Is(err, syscall.ECONNREFUSED)
	}
}
