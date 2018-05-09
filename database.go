package gotools

import (
	"github.com/boycehuang/go-tools/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    //
	_ "github.com/jinzhu/gorm/dialects/mysql"    //
	_ "github.com/jinzhu/gorm/dialects/postgres" //
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //
)

// OpenDB get db connections
func OpenDB(config database.ConnectionConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(config.Dialect, config.URI())
	return
}
