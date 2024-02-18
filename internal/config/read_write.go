package config

import (
	"encoding/json"
	"os"

	"github.com/kaiquegarcia/gojooble/internal/constants"
)

func readFile() (Configuration, error) {
	bytes, err := os.ReadFile(constants.ConfigPath)
	if err != nil {
		return nil, err
	}

	var c configuration
	err = json.Unmarshal(bytes, &c)
	return &c, err
}

func writeFile(c Configuration) error {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(constants.ConfigPath, bytes, os.ModePerm)
}
