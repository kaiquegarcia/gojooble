package config

func Save(c Configuration) error {
	return writeFile(c)
}
