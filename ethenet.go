package main

const ETHER_TYPE_IP uint16 = 0x0800
const ETHER_TYPE_ARP uint16 = 0x0806

var ETHERNET_ADDRESS_BROADCAST = [6]uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

type ethernetHeader struct {
	destAddr  [6]uint8 // XX:XX:XX:XX:XX
	srcAddr   [6]uint8 // XX:XX:XX:XX:XX
	etherType uint16
}

func setMacAddr(macAddrByte []byte) [6]uint8 {
	var macAddrUint8 [6]uint8
	for i, v := range macAddrByte {
		macAddrUint8[i] = v
	}
	return macAddrUint8
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
