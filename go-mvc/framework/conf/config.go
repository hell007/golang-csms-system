package conf

import (
	"github.com/kataras/golog"
	"go-mvc/framework/utils/files"
	"gopkg.in/yaml.v2"
	"runtime"
	"time"
)

type Config struct {
	// 设置
	Charset string `yaml:"Charset"`

	// mysql
	MysqlDialect        string `yaml:"Mysql.Dialect"`
	MysqlPrefix         string `yaml:"Mysql.Prefix"`
	MysqlLogLevel       string `yaml:"Mysql.LogLevel"`
	MysqlMaxIdleConns   int    `yaml:"Mysql.MaxIdleConns"`
	MysqlMaxOpenConns   int    `yaml:"Mysql.MaxOpenConns"`
	MysqlShowSql        bool   `yaml:"Mysql.ShowSql"`
	MysqlMasterUser     string `yaml:"Mysql.Master.User"`
	MysqlMasterPassword string `yaml:"Mysql.Master.Password"`
	MysqlMasterHost     string `yaml:"Mysql.Master.Host"`
	MysqlMasterPort     int    `yaml:"Mysql.Master.Port"`
	MysqlMasterDatabase string `yaml:"Mysql.Master.Database"`
	MysqlSlaveUser      string `yaml:"Mysql.Slave.User"`
	MysqlSlavePassword  string `yaml:"Mysql.Slave.Password"`
	MysqlSlaveHost      string `yaml:"Mysql.Slave.Host"`
	MysqlSlavePort      int    `yaml:"Mysql.Slave.Port"`
	MysqlSlaveDatabase  string `yaml:"Mysql.Slave.Database"`

	// redis
	RedisPrefix          string   `yaml:"Redis.Prefix"`
	RedisExprire         int      `yaml:"Redis.Exprire"`
	RedisSingleAddr      string   `yaml:"Redis.Single.Addr"`
	RedisSinglePassword  string   `yaml:"Redis.Single.Password"`
	RedisSingleDb        int      `yaml:"Redis.Single.Db"`
	RedisSinglePoolSize  int      `yaml:"Redis.Single.PoolSize"`
	RedisClusterAddrs    []string `yaml:"Redis.Cluster.Addrs "`
	RedisClusterPassword string   `yaml:"Redis.Cluster.Password"`
	RedisClusterState    bool     `yaml:"Redis.Cluster.State"`

	// Logs
	LogsWebPath              string `yaml:"Logs.WebPath"`
	LogsAppPath              string `yaml:"Logs.AppPath"`
	LogsType                 string `yaml:"Logs.Type"`
	LogsFileNameDateFormat   string `yaml:"Logs.LogsFileNameDateFormat"`
	LogsTimestampFormat      string `yaml:"Logs.LogsTimestampFormat"`
	LogsLevel                string `yaml:"Logs.Level"`
	LogsOut                  bool   `yaml:"Logs.Out"`
	LogsMaxAge               int    `yaml:"Logs.MaxAge"`
	LogsRotationTime         int    `yaml:"Logs.RotationTime"`
	LogsJSONPrettyPrint      bool   `yaml:"Logs.JSONPrettyPrint"`
	LogsJSONDataKey          string `yaml:"Logs.JSONDataKey"`
	LogsEnableRecordFileInfo bool   `yaml:"Logs.EnableRecordFileInfo"`
	LogsFileInfoField        string `yaml:"Logs.FileInfoField"`

	// jwt
	JWTSecret  string `yaml:"JWT.Secret"`
	JWTSalt    string `yaml:"JWT.Salt"`
	JWTTimeout int    `yaml:"JWT.Timeout"`

	// Auth
	AuthIgnoresWeb []string `yaml:"Auth.Ignores.Web"`
	AuthIgnoresApp []string `yaml:"Auth.Ignores.App"`
	AuthVerifyApp  []string `yaml:"Auth.Verify.App"`
	AuthToken      string   `yaml:"Auth.Token"`

	// Captcha
	CaptchaVerify  bool   `yaml:"Captcha.Verify"`
	CaptchaKey     string `yaml:"Captcha.Key"`
	CaptchaExprire int    `yaml:"Captcha.Exprire"`

	// Upload
	UploadUrl     string   `yaml:"Upload.Url"`
	UploadStyle   []string `yaml:"Upload.Style"`
	UploadPicPath []string `yaml:"Upload.PicPath"`
	UploadEditor  []string `yaml:"Upload.Editor"`

	// time
	Timeformat      string `yaml:"Timeformat"`
	TimeformatShort string `yaml:"TimeformatShort"`
}

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing") // 中国时区
var GlobalConfig Config

/**
获取yaml配置文件
*/
func (cfg *Config) getConf() *Config {
	yamlFile, err := files.LoadFile(GetConfigPath() + "app.yaml")
	if err != nil {
		golog.Errorf("LoadFile cfg config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, cfg); err != nil {
		golog.Errorf("Unmarshal cfg config error!! %s", err)
	}
	return cfg
}

/**
获取配置路径
*/
func GetConfigPath() string {
	var ConfigPath string
	if isWindow() {
		ConfigPath = "D:/Dev/cygwin/work/golang/golang-csms-system/go-mvc/framework/conf/"
	} else {
		ConfigPath = "/Users/wzh/Development/git/go/golang-csms-system/go-mvc/framework/conf/"
	}
	return ConfigPath
}

/**
返回上传文件路径
*/
func GetUploadFile() string {
	var UploadFile string
	if isWindow() {
		UploadFile = "D:/Dev/cygwin/work/golang/golang-csms-system/go-mvc/uploads/"
	} else {
		UploadFile = "/Users/wzh/Development/git/go/golang-csms-system/go-mvc/uploads/"
	}
	return UploadFile
}

func isWindow() bool {
	return runtime.GOOS == "windows"
}

func init() {
	GlobalConfig.getConf()
	golog.Info("配置文件已经加载完成...")
}
