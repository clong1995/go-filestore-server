package config

const (
	// MYSQLSource      : 要连接的数据源
	// root:123456      : 用户名+密码
	// 127.0.0.1:3306   : ip及端口
	// fileserver       : 数据库名
	// charset=utf8     : 使用utf8字符编码
	MYSQLSource = "root:123456@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
)
