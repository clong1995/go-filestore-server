package util

func GetMysqlSource(username, password, host, port, db, charset string) string {
	// "root:123456@tcp(127.0.0.1:3306)/fileServer?charset=utf8"
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=" + charset
}
