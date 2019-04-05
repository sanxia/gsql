package gsql

/* ================================================================================
 * Database and Table Config
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库配置域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DatabaseConfigList []DatabaseConfig
type DatabaseConfig struct {
	TableConfigs TableConfigList
	Source       DatabaseSource
	Name         string
	Count        int
	IsNew        bool
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据表配置域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TableConfigList []TableConfig
type TableConfig struct {
	Name  string
	Count int
	IsNew bool
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据源配置域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DatabaseSource struct {
	Username string
	Password string
	Host     string
}
