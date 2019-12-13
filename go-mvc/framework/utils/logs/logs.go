package logs

import (
	"github.com/go-yaml/yaml"
	"github.com/kataras/golog"

	"go-mvc/framework/conf"
	"go-mvc/framework/utils/files"
)

var (
	Logger MyLog
)

type MyLog  struct {
	Logs LogInfo
}

type LogInfo struct {
	Level  string `yaml:"level"`
	Prefix string `yaml:"prefix"`
	Output string `yaml:"output"`
}

func (log *MyLog) GetConf() *MyLog {
	yamlFile, err := files.LoadFile(conf.GetConfigPath() + "conf.yaml")
	if err != nil {
		golog.Errorf("LoadFile logs config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, log); err != nil {
		golog.Errorf("Unmarshal logs config error!! %s", err)
	}
	return log
}
