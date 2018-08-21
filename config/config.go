package config

import (
	"encoding/json"
	"io/ioutil"
)

type Statement struct {
	Effect   string
	Action   []string
	Resource string
}

type Config struct {
	Name     string
	VCSURL   string
	Profile  string
	VpcID    string
	Subnets  []string
	Policies []Statement
}

func (conf Config) ProjectName() string {
	return conf.Profile + ":" + conf.Name
}

func Read(path string) (Config, error) {
	var conf Config

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(bytes, &conf)

	return conf, err
}

func Write(path string, config Config) error {
	s, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return writeFile(path, s)
}

func writeFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}
