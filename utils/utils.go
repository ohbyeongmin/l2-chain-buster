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

func ConvertStructsToYAML(filename string, any ...interface{}) error {
	var data []byte
	for _, s := range any {
		yamlData, err := yaml.Marshal(s)
		if err != nil {
			return err
		}
		data = append(data, yamlData...)
	}
	os.WriteFile(filename, data, 0644)
	return nil
}
