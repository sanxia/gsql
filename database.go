package gsql

import (
	"log"
	"time"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/* ================================================================================
 * Database
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DatabaseList []*Database
type Database struct {
	db           *sql.DB
	Source       DatabaseSource
	tables       TableList
	Name         string
	shardingName string
	Sql          string
	IsNew        bool
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据表域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TableList []*Table
type Table struct {
	Database     *Database
	Name         string
	shardingName string
	Sql          string
	IsNew        bool
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 打开数据库连接
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Database) Open(args ...string) error {
	database := ""
	if len(args) > 0 {
		database = args[0]
	}

	dsn := s.Source.Username + ":" + s.Source.Password + "@tcp(" + s.Source.Host + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("sql Open error: %v", err)
		return err
	}

	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxOpenConns(128)
	db.SetMaxIdleConns(16)

	s.db = db

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 运行Sql语句
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Database) Exec(sql string, args ...string) (sql.Result, error) {
	if s.db == nil {
		database := ""
		if len(args) > 0 {
			database = args[0]
		}

		if err := s.Open(database); err != nil {
			log.Printf("sql Open error: %v", err)
			return nil, err
		}
	}

	sqlResult, err := s.db.Exec(sql)
	if err != nil {
		log.Printf("sqlDb.Exec error: %v", err)
		return nil, err
	}

	return sqlResult, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sql集合查询
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Database) Query(sql string, args ...string) ([]map[string]interface{}, error) {
	if s.db == nil {
		database := ""
		if len(args) > 0 {
			database = args[0]
		}

		if err := s.Open(database); err != nil {
			log.Printf("sql Open error: %v", err)
			return nil, err
		}
	}

	records := make([]map[string]interface{}, 0)
	rows, err := s.db.Query(sql)
	if err != nil {
		log.Printf("sqlDb.Query error: %v", err)
		return nil, err
	}

	for rows.Next() {
		record := make(map[string]interface{})
		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		scanValues := make([]interface{}, len(columns))

		for i := range scanValues {
			scanArgs[i] = &scanValues[i]
		}

		err = rows.Scan(scanArgs...)
		for i, value := range scanValues {
			if value != nil {
				record[columns[i]] = value
			}
		}

		records = append(records, record)
	}

	rows.Close()

	return records, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 关闭数据库连接
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Database) Close() error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			log.Printf("sql Close error: %v", err)

			return err
		}

		s.db = nil
	}

	return nil
}
