package lib

import (
	"fmt"
	"os"
	"strconv"
)

type Configuration interface {
	// constructor
	NewConfig() Config

	// @getter: Config.Port
	GetPort() uint16
	// @setter: Config.Port
	SetPort(port uint16)
	// @logger: Config.Port
	LogginPort()

	// @setter: ApiKeys (from .env file in the root of the project: server folder directory).
	SetApiKeys()
	// @getter: ApiKeys (from .env file in the root of the project: server folder directory).
	GetApiKeys(key string) string
}

type Config struct {
	// Port number of server.
	Port uint16
	// All APIs authorization keys for 3rd party microservices.
	ApiKeys map[string]string
}

func NewConfig() Config {
	/*
		@description:
			- Constructor of config struct.
			- It will be accessable from Configure interface.
		i dunno, but i use that port
		for a port num problem.
	*/
	value, err := strconv.ParseInt(os.Getenv("GO_PORT"), 10, 16)
	if err != nil {
		value = 4567
	}

	return Config{Port: uint16(value), ApiKeys: nil}
}

func (c Config) LoggingPort() {
	/*
		@description:
			- It'll be print the current port number of golang server.
	*/
	fmt.Println("Current port is: ", c.Port)
}

func (c *Config) SetPort(port uint16) {
	/*
		@description:
			- Setter of Port number in Config struct.

		@params:
			- port: int16 = wanted port number of project server to serve.
	*/
	c.Port = port
}

func (c *Config) GetPort() uint16 {
	/*
		@description:
			- Getter of Port number in Config struct.
	*/
	return c.Port
}

func (c *Config) SetApiKeys(key_name, value_name string) {
	/*
		@description:
			- Getter of API Keys map from .env file in the root of golang project.
	*/
	// setting
	c.ApiKeys[key_name] = value_name
}

func (c Config) GetApiKeys(key string) string {
	/*
		@description:
			- Getter for Api Key (map[string]string).
			-- The keys are stored in .env files
	*/
	return c.ApiKeys[key]
}
