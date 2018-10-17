package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	}
	Cache struct {
		Enable bool `yaml:"enable"`
		List []string `yaml:"list,flow"`
	}
}


func main() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("test.yaml")

	fmt.Println("yamlFile content is:", string(yamlFile))
	if err != nil {
		fmt.Println("yamlFile.Get err ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		fmt.Println("err is:", err)
	}
	fmt.Println("conf", conf)
	fmt.Println("conf.mysql is",conf.Mysql)
	fmt.Println("conf.cache is",conf.Cache)

}
