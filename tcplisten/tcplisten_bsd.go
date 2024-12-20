//go:build darwin || dragonfly || freebsd || netbsd || openbsd || rumprun

package tcplisten

import (
	"fmt"
	"syscall"
)

const soReusePort = syscall.SO_REUSEPORT

func enableDeferAccept(fd int) error {
	// TODO: implement SO_ACCEPTFILTER:dataready here
	return fmt.Errorf("doesn't support SO_ACCEPTFILTER")
}

func enableFastOpen(fd int) error {
	// TODO: implement TCP_FASTOPEN when it will be ready
	return fmt.Errorf("doesn't support TCP_FASTOPEN")
}

func soMaxConn() (int, error) {
	// TODO: properly implement it
	return syscall.SOMAXCONN, nil
}
