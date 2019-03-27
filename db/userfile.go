package db

import (
	"fmt"
	mydb "go-filestore-server/db/mysql"
	"time"
)

// UserFile: 用户文件表结构体
type UserFile struct {
	UserName    string `json:"user_name"`    // 用户名
	FileHash    string `json:"file_hash"`    // 文件Hash
	FileName    string `json:"file_name"`    // 文件名
	FileSize    int64  `json:"file_size"`    // 文件尺寸
	UploadAt    string `json:"upload_at"`    // 上传时间
	LastUpdated string `json:"last_updated"` // 最后一次更新时间
}

// 添加: 插入用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	// insert ignore 会忽略数据库中已经存在的数据
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user_file (`user_name`,`file_hash`,`file_name`,`file_size`,`upload_at`) values(?,?,?,?,?) ")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, time.Now())
	if err != nil {
		return false
	}
	return true
}

// 更新: 文件重命名
func RenameFileName(username, filehash, filename string) bool {
	stmt, err := mydb.DBConn().Prepare("update tbl_user_file set file_name = ? where user_name = ? and file_hash = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(filename, username, filehash)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// 查询: 批量获取用户文件信息
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_hash, file_name, file_size, upload_at,last_update from tbl_user_file where user_name=? and limit ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil {
		return nil, err
	}

	var userFiles []UserFile
	for rows.Next() {
		ufile := UserFile{}
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSize, &ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		userFiles = append(userFiles, ufile)
	}
	return userFiles, nil
}

// 查询: 用户单个文件信息
func QueryUserFileMete(username string, filehash string) (*UserFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_hash, file_name, file_size, upload_at, last_update from tbl_user_file where user_name = ? and file_hash =? limit 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, filehash)
	if err != nil {
		return nil, err
	}

	ufile := UserFile{}
	if rows.Next() {
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSize, &ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return &ufile, nil
}
