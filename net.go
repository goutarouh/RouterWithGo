package main

import (
	"fmt"
	"syscall"
)

// ルータが持つNWインターフェースを表す
type netDevice struct {
	name       string
	macaddr    [6]uint8
	socket     int // NWインターフェースにbindするソケットのディスクリプタ番号
	sockaddr   syscall.SockaddrLinklayer
	etheHeader ethernetHeader
	ipdev      ipDevice
}

var IGNORE_INTERFACES = []string{"lo", "bond0", "dummy0", "tunl0", "sit0"}

func isIgnoreInterfaces(name string) bool {
	for _, v := range IGNORE_INTERFACES {
		if v == name {
			return true
		}
	}
	return false
}

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func (netdev *netDevice) netDevicePoll(mode string) error {
	recvbuffer := make([]byte, 1500)
	n, _, err := syscall.Recvfrom(netdev.socket, recvbuffer, 0)
	if err != nil {
		if n == 01 {
			return nil
		} else {
			return fmt.Errorf("recv err, n is %d, device is %s, err is %s", n, netdev.name, err)
		}
	}
	if mode == "ch1" {
		fmt.Printf("received %d bytes from %s: %x\n", n, netdev.name, recvbuffer[:n])
	} else {
		ethernetInput(netdev, recvbuffer[:n])
	}
	return nil
}
