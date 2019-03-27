package db

import (
	"database/sql"
	"fmt"
	mydb "go-filestore-server/db/mysql"
)

// 插入:  文件上传完成，保存meta
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_file (`file_hash`,`file_name`,`file_size`,`file_addr`,`status`) value (?,?,?,?,1)")
	if err != nil {
		fmt.Println("failed to prepare statement, err:\t", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Println("file with hash has been uploaded before:\t", filehash)
		}
		return true
	}
	return false
}

// TableFile: 文件表结构体
type TableFile struct {
	FileHash string         `json:"file_hash"`
	FileName sql.NullString `json:"file_name"`
	FileSize sql.NullInt64  `json:"file_size"`
	FileAddr sql.NullString `json:"file_addr"`
}

// 查询:  获取文件云信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_hash, file_addr, file_name, file_size from tbl_file where file_hash = ? and status = 1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		if err == sql.ErrNoRows {
			// 查不到对应记录，返回参数及错误均为nil
			return nil, nil
		} else {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return &tfile, nil
}

// 查询:  批量获取文件元信息
func GetFileMetaList(limit int) ([]TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_hash, file_addr, file_name, file_size from tbl_file where status=1 limit ?")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	cloumns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(cloumns))
	var tFiles []TableFile
	for i := 0; i < len(values) && rows.Next(); i++ {
		tFile := TableFile{}
		err = rows.Scan(&tFile.FileHash, &tFile.FileAddr, &tFile.FileName, &tFile.FileSize)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		tFiles = append(tFiles, tFile)
	}
	fmt.Println(len(tFiles))
	return tFiles, nil
}

// 更新:  更新文件的存储地址
func UpdateFileLocation(filehash string, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare("update tbl_file set `file_addr` = ? where `file_hash` = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(fileaddr, filehash)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Println("更新文件addr 失败, filehash:\t", filehash)
		}
		return true
	}
	return false
}
