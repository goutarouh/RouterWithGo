package main

type radixTreeNode struct {
	depth  int
	parent *radixTreeNode
	node0  *radixTreeNode
	node1  *radixTreeNode
	data   ipRouteEntry
	value  int
}

func (node *radixTreeNode) radixTreeAdd(prefixIpAddr, prefixLen uint32, entryData ipRouteEntry) {
	current := node
	for i := 1; i <= int(prefixLen); i++ {
		if prefixIpAddr>>(32-i)&0x01 == 1 { // 上からiビット目が1なら
			if current.node1 == nil {
				current.node1 = &radixTreeNode{
					parent: current,
					depth:  i,
					value:  0,
				}
			}
			current = current.node1
		} else {
			if current.node0 == nil {
				current.node0 = &radixTreeNode{
					parent: current,
					depth:  i,
					value:  0,
				}
			}
			current = current.node0
		}
	}
	current.data = entryData
}

func (node *radixTreeNode) radixTreeSearch(prefixIpAddr uint32) ipRouteEntry {
	current := node
	var result ipRouteEntry
	// 検索するIPアドレスと比較して1ビットずつ辿っていく
	for i := 1; i <= 32; i++ {
		if current.data != (ipRouteEntry{}) {
			result = current.data
		}
		if (prefixIpAddr>>(32-i))&0x01 == 1 { // 上からiビット目が1だったら
			if current.node1 == nil {
				return result
			}
			current = current.node1
		} else { // iビット目が0だったら
			if current.node0 == nil {
				return result
			}
			current = current.node0
		}
	}
	return result
}
