package main

import "fmt"

// 存储类型（表示文件存到哪里)
type StoreType int

const (
	_          StoreType = iota
	StoreLocal           // 节点本地
	StoreCeph            // Ceph集群
	StoreOSS             // 阿里OSS
	StoreMix             // 混合(Ceph+OSS)
	StoreAll             // 所有类型的存储都存储一份数据
)

func main() {
	fmt.Println(StoreLocal)
	fmt.Println(StoreCeph)
	fmt.Println(StoreOSS)
	fmt.Println(StoreMix)
	fmt.Println(StoreAll)
}
