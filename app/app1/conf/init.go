package conf

import "github.com/pengzhong2010/go-server-base/common/pkg/util"

var (
	APP_PORT      string = util.GetEnv("XXX_APP_PORT", "8080")
	MYSQL_HOST    string = util.GetEnv("XXX_MYSQL_HOST", "192.168.0.50")
	MYSQL_PORT    string = util.GetEnv("XXX_MYSQL_PORT", "58001")
	MYSQL_USER    string = util.GetEnv("XXX_MYSQL_USER", "root")
	MYSQL_PASSWD  string = util.GetEnv("XXX_MYSQL_PASSWD", "caidaoninb")
	MYSQL_DB_NAME string = util.GetEnv("XXX_MYSQL_DB_NAME", "trolley")
	ES_URL        string = util.GetEnv("XXX_ES_URL", "192.168.1.1,192.168.1.2,192.168.1.3")
)
