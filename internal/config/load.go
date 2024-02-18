package config

import "os"

func Load() (Configuration, error) {
	c, err := readFile()
	if os.IsNotExist(err) {
		return New(""), nil
	}

	return c, err
}
