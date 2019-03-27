package mq

import "go-filestore-server/common"

// 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash      string           `json:"file_hash"`
	CurLocation   string           `json:"cur_location"`
	DestLocation  string           `json:"dest_location"`
	DestStoreType common.StoreType `json:"dest_store_type"`
}
