package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func sumByteArr(packet []byte) (sum uint) {
	for i, _ := range packet {
		if i%2 == 0 {
			sum += uint(byteToUint16(packet[i:]))
		}
	}
	return sum
}

func calcChecksum(packet []byte) []byte {
	// まず16ビット毎に足す
	sum := sumByteArr(packet)
	// あふれた桁を足す
	sum = (sum & 0xffff) + sum>>16
	// 論理否定を取った値をbyteにして返す
	return uint16ToByte(uint16(sum ^ 0xffff))
}

func printMacAddr(macddr [6]uint8) string {
	var str string
	for _, v := range macddr {
		str += fmt.Sprintf("%x:", v)
	}
	return strings.TrimRight(str, ":")
}

func byteToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func byteToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func uint16ToByte(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func uint32ToByte(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

/*
intを10進数のipアドレス表記に変換
*/
func convertuint32ToIPv4(ipInt uint32) string {
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt >> 24)
	ip[1] = byte(ipInt >> 16)
	ip[2] = byte(ipInt >> 8)
	ip[3] = byte(ipInt)

	return ip.String()
}
