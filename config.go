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
	Name         string
	Count        int
	TableConfigs TableConfigList
	Source       DatabaseSource
	IsNew        bool
	IsSharding   bool //是否分片
	IsBuilded    bool //是否已生成过
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据表配置域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type TableConfigList []TableConfig
type TableConfig struct {
	Name       string
	Count      int
	IsNew      bool
	IsSharding bool //是否分片
	IsBuilded  bool //是否已生成过
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据源配置域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DatabaseSource struct {
	Username string
	Password string
	Host     string
}
