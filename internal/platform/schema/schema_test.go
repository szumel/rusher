package schema

import (
	"reflect"
	"testing"
)

func TestGetCurrentConfig(t *testing.T) {
	c := Config{Environment: "test"}
	s := ConfigPool{
		Configs: []Config{c},
	}
	env := "test"

	expected, err := GetCurrentConfig(&s, env)
	if err != nil {
		t.Fatal("Could not get current config.")
	}

	if !reflect.DeepEqual(&c, expected) {
		t.Fatalf("Could not get current config. Expected %+v have %+v.", c, expected)
	}
}
