package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ConvertYAMLtoStruct(filename string, any interface{}) error {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlFile, any); err != nil {
		return err
	}
	return nil
}
