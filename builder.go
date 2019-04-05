package gsql

import (
	"fmt"
	"log"
	"strings"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * sqlBuilder
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 索引选项数据域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 分片数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type sqlBuilder struct {
	Configs   DatabaseConfigList
	databases DatabaseList
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化SqlBuilder
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewSqlBuilder() *sqlBuilder {
	builder := &sqlBuilder{
		databases: make(DatabaseList, 0),
	}

	return builder
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化配置
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *sqlBuilder) config() {
	for _, databaseConfig := range s.Configs {
		for i := 0; i < databaseConfig.Count; i++ {
			database := new(Database)
			database.Source = databaseConfig.Source
			database.Name = databaseConfig.Name

			if databaseConfig.Count > 1 {
				database.shardingName = fmt.Sprintf("%s_%05d", database.Name, i)
			} else {
				database.shardingName = database.Name
			}

			database.IsNew = databaseConfig.IsNew

			database.tables = make(TableList, 0)

			for _, tableConfig := range databaseConfig.TableConfigs {
				for j := 0; j < tableConfig.Count; j++ {
					table := new(Table)
					table.Name = tableConfig.Name

					if tableConfig.Count > 1 {
						table.shardingName = fmt.Sprintf("%s_%05d", table.Name, j)

					} else {
						table.shardingName = table.Name
					}

					table.Database = database
					table.IsNew = tableConfig.IsNew

					database.tables = append(database.tables, table)
				}
			}

			s.databases = append(s.databases, database)
		}
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *  生成
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *sqlBuilder) Build() {
	s.config()

	for _, database := range s.databases {
		s.createDatabase(database)

		for _, table := range database.tables {
			s.createTable(table)
		}

		log.Printf("%s", "\r\n")
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建数据库
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *sqlBuilder) createDatabase(database *Database) {
	defer database.Close()

	if database.IsNew {
		sql := fmt.Sprintf("DROP DATABASE IF EXISTS `%s`;", database.shardingName)
		if _, err := database.Exec(sql); err != nil {
			log.Printf("sqlDb.Exec error: %v", err)
		}
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", database.shardingName)
	if _, err := database.Exec(sql); err != nil {
		log.Printf("sqlDb.Exec error: %v", err)
	}

	log.Printf("databaseName: %s", database.shardingName)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建数据表
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *sqlBuilder) createTable(table *Table) {
	defer table.Database.Close()

	if table.Database.IsNew || table.IsNew {
		sql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", table.shardingName)
		if _, err := table.Database.Exec(sql, table.Database.shardingName); err != nil {
			log.Printf("sqlDb.Exec error: %v", err)
		}

		if err := s.getTableSql(table); err != nil {
			log.Printf("getTableSql error: %v", err)
		} else {
			if _, err := table.Database.Exec(table.Sql, table.Database.shardingName); err != nil {
				log.Printf("sqlDb.Exec error: %v", err)
			}
		}
	}

	log.Printf("tableName: %s", table.shardingName)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据表Sql
 * Sql优先，否则从文件获取
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *sqlBuilder) getTableSql(table *Table) error {
	if len(table.Sql) == 0 {
		filename := fmt.Sprintf("%s/sql/%s/%s.sql", glib.GetCurrentPath(), table.Database.Name, table.Name)
		content, err := glib.GetFileContent(filename)
		if err != nil {
			return err
		}

		table.Sql = string(content)
	}

	table.Sql = strings.Replace(table.Sql, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", -1)
	table.Sql = strings.Replace(table.Sql, table.Name, table.shardingName, -1)

	return nil
}
