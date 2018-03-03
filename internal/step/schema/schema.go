//package schema represents configuration schema which contains context and jobs information
package schema

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"reflect"
)

const (
	errInvalidEnv = "ensure that given env exists in config pool"
)

func New(path string) (*ConfigPool, error) {
	c := &ConfigPool{}
	err := c.Fetch(path)

	return c, err
}

type ConfigPool struct {
	XMLName xml.Name `xml:"configPool"`
	Configs []Config `xml:"config"`
}

type Config struct {
	Environment string   `xml:"environment"`
	XMLName     xml.Name `xml:"config"`
	ProjectPath string   `xml:"projectPath"`
	Steps       []Step   `xml:"steps>step"`
}

type Step struct {
	Code   string     `xml:"code,attr"`
	Params []xml.Attr `xml:",any,attr"`
}

func (c *ConfigPool) Fetch(filename string) error {
	configData, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = xml.Unmarshal(configData, c)

	if err != nil {
		return err
	}

	return nil
}

func GetCurrentConfig(s *ConfigPool, env string) (*Config, error) {
	var currentConfig Config
	for _, config := range s.Configs {
		if config.Environment == env {
			currentConfig = config
			break
		}
	}

	if reflect.DeepEqual(Config{}, currentConfig) {
		return &Config{}, errors.New(errInvalidEnv)
	}

	return &currentConfig, nil
}
