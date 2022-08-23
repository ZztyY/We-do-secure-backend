package database

import (
	"github.com/jinzhu/gorm"
)

func ExecSQL(sql string) *gorm.DB {
	return DB.Exec(sql)
}

func RowSQLCount(sql string) int {
	var count int
	DB.Raw(sql).Count(&count)
	return count
}
