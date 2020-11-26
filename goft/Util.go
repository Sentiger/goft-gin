package goft

import (
	"io/ioutil"
	"log"
	"os"
)

// 读取配置文件
func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	return b
}
