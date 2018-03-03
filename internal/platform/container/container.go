package container

import "errors"

const errInvalidService = "service does not exist for given serviceName"

var services map[string]interface{}

func init() {
	services = map[string]interface{}{}
}

func Get(serviceName string) (interface{}, error) {
	s, ok := services[serviceName]
	if !ok {
		return nil, errors.New(errInvalidService)
	}

	return s, nil
}

func Set(serviceName string, instance interface{}) error {
	services[serviceName] = instance
	//@todo check if duplicated?
	return nil
}


