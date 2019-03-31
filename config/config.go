package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Env                  string `json:"env"`                      // 环境
	UploadServiceHost    string `json:"upload_service_host"`      // 上传服务监听的地址
	PwdSalt              string `json:"pwd_salt"`                 // 盐值
	CephAccessKey        string `json:"ceph_access_key"`          // Ceph访问Key
	CephSecretKey        string `json:"ceph_secret_key"`          // Ceph访问密钥
	CephGWEndpoint       string `json:"ceph_gw_endpoint"`         // Ceph Gateway地址
	MysqlUser            string `json:"mysql_user"`               // mysql用户名
	MysqlPwd             string `json:"mysql_pwd"`                // mysql密码
	MysqlHost            string `json:"mysql_host"`               // mysql ip
	MysqlPort            string `json:"mysql_port"`               // mysql port
	MysqlDb              string `json:"mysql_db"`                 // mysql db
	MysqlCharset         string `json:"mysql_charset"`            // mysql charset
	MysqlMaxConn         int    `json:"mysql_max_conn"`           // mysql max connect
	RedisHost            string `json:"redis_host"`               // redis地址
	RedisPass            string `json:"redis_pass"`               // redis密码
	OSSBucket            string `json:"oss_bucket"`               // oss bucket名
	OSSEndpoint          string `json:"oss_endpoint"`             // oss endpoint
	OSSAccessKey         string `json:"oss_access_key"`           // oss 访问key
	OSSAccessSecret      string `json:"oss_access_secret"`        // oss 访问secret
	AsyncTransferEnable  bool   `json:"async_transfer_enable"`    // 是否开启文件异步转移(默认同步)
	RabbitURL            string `json:"rabbit_url"`               // rabbitmq服务的入口url
	TransExchangeName    string `json:"trans_exchange_name"`      // 用户文件transfer的交换机
	TransOSSQueueName    string `json:"trans_oss_queue_name"`     // oss转移队列名
	TransOSSErrQueueName string `json:"trans_oss_err_queue_name"` // oss转移失败后写入另一个队列的队列名
	TransOSSRoutingKey   string `json:"trans_oss_routing_key"`    // routingkey
	TempLocalRootDir     string `json:"temp_local_root_dir"`      // 本地临时存储地址路径
	CurrentStoreType     int    `json:"current_store_type"`       // 设置当前文件的存储类型
	OSSRootDir           string `json:"oss_root_dir"`             // OSS的存储路径前缀
	CephRootDir          string `json:"ceph_root_dir"`            // Ceph的存储路径前缀
}

var DefaultConfig *Configuration

func InitConfig(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	DefaultConfig = &Configuration{}

	err := decoder.Decode(DefaultConfig)
	if err != nil {
		panic(err)
	}
}
