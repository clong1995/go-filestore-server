package common

// 存储类型（表示文件存到哪里)
type StoreType int

const (
	_          StoreType = iota
	StoreLocal           // 1.节点本地
	StoreCeph            // 2.Ceph集群
	StoreOSS             // 3.阿里OSS
	StoreMix             // 4.混合(Ceph+OSS)
	StoreAll             // 5.所有类型的存储都存储一份数据
)
