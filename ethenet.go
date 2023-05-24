package main

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
