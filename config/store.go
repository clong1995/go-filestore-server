package config

import "go-filestore-server/common"

const (
	TempLocalRootDir = "/data/fileserver" // 本地临时存储地址路径
	CurrentStoreType = common.StoreLocal  // 设置当前文件的存储类型
)
