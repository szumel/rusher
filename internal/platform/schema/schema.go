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

func NewFromString(schema string) (*ConfigPool, error) {
	return NewFromByte([]byte(schema))
}

func NewFromByte(schema []byte) (*ConfigPool, error) {
	c := &ConfigPool{}
	err := xml.Unmarshal(schema, c)

	return c, err
}

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
	Sequence    Sequence `xml:"sequence"`
	Globals     []Var    `xml:"globals>var"`
}

type Sequence struct {
	SequenceElems []SequenceElem `xml:",any"`
}

type Var struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

//@todo create StepElem like MacroElem and convert (remove Code there)
type SequenceElem struct {
	XMLName xml.Name
	Code    string     `xml:"code,attr"`
	Params  []xml.Attr `xml:",any,attr"`
}

type MacroElem struct {
	Version string
	Source  string
	Params  []xml.Attr
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
