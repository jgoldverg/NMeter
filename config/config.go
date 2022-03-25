package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	File struct {
		DataFileName string `yaml:"dataFileName"`
		DataFilePath string `yaml:"dataFilePath"`
	} `yaml:"config"`
}

//The config file will always be in ~/.pmeter but we want to know where to store the measurements.
func ReadConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}

func prepareDir(filePath string) {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteConfig(config *Config, filePath string) {
	prepareDir(filePath)
	data, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := ioutil.WriteFile(filePath, data, 0)
	if err2 != nil {
		fmt.Printf(err2.Error())
	}
	fmt.Println("wrote config file")
}
