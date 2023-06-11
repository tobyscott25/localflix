package helper

import (
	"os"

	"gopkg.in/yaml.v3"
)

func WriteYAMLFile(filename string, data interface{}) error {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(yamlBytes)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
