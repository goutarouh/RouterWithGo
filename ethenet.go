package main

import (
	"bytes"
	"log"
)

const ETHER_TYPE_IP uint16 = 0x0800
const ETHER_TYPE_ARP uint16 = 0x0806
const ETHERNET_ADDRES_LEN = 6

var ETHERNET_ADDRESS_BROADCAST = [6]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

type ethernetHeader struct {
	destAddr  [6]uint8 // XX:XX:XX:XX:XX
	srcAddr   [6]uint8 // XX:XX:XX:XX:XX
	etherType uint16
}

func (ethHeader ethernetHeader) ToPacket() []byte {
	var b bytes.Buffer
	b.Write(macToByte(ethHeader.destAddr))
	b.Write(macToByte(ethHeader.srcAddr))
	b.Write(uint16ToByte(ethHeader.etherType))
	return b.Bytes()
}

func setMacAddr(macAddrByte []byte) [6]uint8 {
	var macAddrUint8 [6]uint8
	for i, v := range macAddrByte {
		macAddrUint8[i] = v
	}
	return macAddrUint8
}

func macToByte(macAddr [6]uint8) (b []byte) {
	for _, v := range macAddr {
		b = append(b, v)
	}
	return b
}

func ethernetInput(netdev *netDevice, packet []byte) {
	netdev.etheHeader.destAddr = setMacAddr(packet[0:6])
	netdev.etheHeader.srcAddr = setMacAddr(packet[6:12])
	netdev.etheHeader.etherType = byteToUint16(packet[12:14])

	if netdev.macaddr != netdev.etheHeader.destAddr && netdev.etheHeader.destAddr != ETHERNET_ADDRESS_BROADCAST {
		// 自分が宛先でない場合かつブロードキャストではない場合は処理を終了する
		return
	}

	switch netdev.etheHeader.etherType {
	case ETHER_TYPE_ARP:
		arpInput(netdev, packet[14:])
	case ETHER_TYPE_IP:
		ipInput(netdev, packet[14:])
	}
}

func ethernetOutput(netdev *netDevice, destaddr [6]uint8, packet []byte, ethType uint16) {
	// イーサネットヘッダのパケットを作成
	ethHeaderPacket := ethernetHeader{
		destAddr:  destaddr,
		srcAddr:   netdev.macaddr,
		etherType: ethType,
	}.ToPacket()
	// イーサネットヘッダに送信するパケットをつなげる
	ethHeaderPacket = append(ethHeaderPacket, packet...)
	// ネットワークデバイスに送信する
	err := netdev.netDeviceTransmit(ethHeaderPacket)
	if err != nil {
		log.Fatalf("netDeviceTransmit is err : %v", err)
	}
}
