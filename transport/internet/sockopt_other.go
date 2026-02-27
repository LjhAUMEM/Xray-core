//go:build js || netbsd || openbsd || solaris
// +build js netbsd openbsd solaris

package internet

import "github.com/xtls/xray-core/common/net"

func applyOutboundSocketOptions(network string, address string, fd uintptr, config *SocketConfig) error {
	return nil
}

func applyInboundSocketOptions(network string, fd uintptr, config *SocketConfig) error {
	return nil
}

func bindAddr(fd uintptr, ip []byte, port uint32) error {
	return nil
}

func setReuseAddr(fd uintptr) error {
	return nil
}

func setReusePort(fd uintptr) error {
	return nil
}

func setBrutalRate(conn net.Conn, rate uint64, cwnd uint32) error {
	return nil
}
