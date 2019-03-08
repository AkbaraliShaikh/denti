package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"denti"`
	Port    string `default:"8282"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"denti.log"`
	}
	DB struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"postgres"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"postgres"`
			Database string `default:"denti"`
		}
	}
	Contacts struct {
		Name  string `default:"Akbar Shaikh"`
		Email string `default:"aashaikh55@gmail.com"`
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "pkg/config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
