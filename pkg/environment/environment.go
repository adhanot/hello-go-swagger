package environment

import (
	"os"
	"strconv"
)

const (
	Production  = "prod"
	Development = "dev"

	Default        = "Unknown"
	DefaultVersion = "v0.0.1"

	DownloadFolder = "/tmp/hello"

	Host    = "0.0.0.0"
	RootURL = "http://localhost:8081"

	Port = 8080
)

type Environment struct {
	Version string
	Name    string

	DownloadFolder string

	RootURL string
	Host    string
	Port    int
}

var Current *Environment

func IsProduction() bool {
	if Current != nil && Current.Name == Production {
		return true
	}
	return false
}

func IsDevelopment() bool {
	if Current != nil && Current.Name == Development {
		return true
	}
	return false
}

func GetEnvironment() *Environment {
	if Current != nil {
		return Current
	}
	env := Environment{}
	if v, ok := os.LookupEnv("HELLO_ENV"); ok {
		env.Name = v
	} else {
		env.Name = Default
	}

	if v, ok := os.LookupEnv("HELLO_VERSION"); ok {
		env.Version = v
	} else {
		env.Version = DefaultVersion
	}

	if v, ok := os.LookupEnv("HELLO_DATA"); ok {
		env.DownloadFolder = v
	} else {
		env.DownloadFolder = DownloadFolder
	}

	if v, ok := os.LookupEnv("ROOT_URL"); ok {
		env.RootURL = v
	} else {
		env.RootURL = RootURL
	}
	if v, ok := os.LookupEnv("HOST"); ok {
		env.Host = v
	} else {
		env.Host = Host
	}
	if v, ok := os.LookupEnv("PORT"); ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			env.Port = Port
		}
		env.Port = int(i)
	} else {
		env.Port = Port
	}

	Current = &env
	return &env
}
