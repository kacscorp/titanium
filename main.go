package main

import (
	"flag"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/kacscorp/titanium/lib/config"
)

var (
	configuration *config.Configuration
)

const (
	defaultConfigFilename = "conf.yaml"
)

func init() {
	env := "development"
	flag.StringVar(&env, "env", "", "Set the environment (development, testing, staging or production)")
	flag.Parse()

	if err := readConfigFile(); err != nil {
		panic("failed to load configuration")
	}

}

func main() {
	err := config.Run(configuration)
	if err != nil {
		panic(err)
	}
}

func readConfigFile() error {
	bytes, err := ioutil.ReadFile(defaultConfigFilename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bytes, &configuration); err != nil {
		return err
	}
	return nil
}
