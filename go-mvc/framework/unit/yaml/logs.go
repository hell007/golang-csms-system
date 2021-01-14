package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/kataras/golog"

	"go-mvc/framework/utils/files"
)

type MyLog struct {
	Logs LogInfo
}

type LogInfo struct {
	Level  string `yaml:"level"`
	Prefix string `yaml:"prefix"`
	Output string `yaml:"output"`
}

func (log *MyLog) getConf() *MyLog {
	yamlFile, err := files.LoadFile("./conf.yaml")
	if err != nil {
		golog.Errorf("LoadFile logs config error!! %s", err)
	}

	if err = yaml.Unmarshal(yamlFile, log); err != nil {
		golog.Errorf("Unmarshal logs config error!! %s", err)
	}
	return log
}

func main() {
	var log MyLog
	log.getConf()
	fmt.Println("log===", log.Logs)
	fmt.Println("Output===", log.Logs.Output)
}
