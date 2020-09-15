package conn

import (
	dbm "github.com/CLannadZSY/go-library/database/db"
)

var DB *dbm.DB

// 创建数据库连接
func InitDb(writeName, readName, filePath string) *dbm.DB {
	DB = dbm.ConnectMysql(writeName, readName, filePath)
	return DB
}
