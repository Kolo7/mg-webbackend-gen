package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jimsmart/schema"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kolo7/mg-webbackend-gen/options"
)

var (
	_DB              *sql.DB
	_dbTables        []string
	_dbExcludeTables []string
)

func InitSchema() error {
	db, err := initializeDB()
	if err != nil {
		return err
	}
	_DB = db
	return nil
}

func GetDB() *sql.DB {
	return _DB
}

func GetTableNames() ([]string, error) {
	if len(_dbTables) != 0 {
		return _dbTables, nil
	}
	var dbTables []string
	if *options.SqlTable != "" {
		dbTables = strings.Split(*options.SqlTable, ",")
	} else {
		schemaTables, err := schema.TableNames(_DB)
		if err != nil {
			fmt.Print(options.Au.Red(fmt.Sprintf("Error in fetching tables information from %s information schema from %s\n",
				*options.SqlType, *options.SqlConnStr)))
			return nil, err
		}
		for _, st := range schemaTables {
			dbTables = append(dbTables, st[1]) // s[0] == sqlDatabase
		}
	}
	_dbTables = dbTables
	return dbTables, nil
}

func GetExcludeTableName() ([]string, error) {
	if len(_dbExcludeTables) != 0 {
		return _dbExcludeTables, nil
	}

	var excludeDbTables []string

	if *options.ExcludeSQLTables != "" {
		excludeDbTables = strings.Split(*options.ExcludeSQLTables, ",")
	}
	return excludeDbTables, nil
}

func initializeDB() (db *sql.DB, err error) {
	db, err = sql.Open(*options.SqlType, *options.SqlConnStr)
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error in open database: %v\n\n", err.Error())))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(options.Au.Red(fmt.Sprintf("Error pinging database: %v\n\n", err.Error())))
		return
	}

	return
}
