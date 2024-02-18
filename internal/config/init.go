package config

type Configuration interface {
	ApiKey() string
	IsEmpty() bool
}

func New(apiKey string) Configuration {
	return &configuration{
		IApiKey: apiKey,
	}
}

type configuration struct {
	IApiKey string `json:"api_key"`
}

func (c *configuration) ApiKey() string {
	return c.IApiKey
}

func (c *configuration) IsEmpty() bool {
	return c.IApiKey == ""
}
