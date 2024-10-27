package configs

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type ProgramConfig struct {
	Server    int
	DBPort    int
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	Secret    string
	RefSecret string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res, errRes := loadConfig()

	logrus.Error("Error Load Config : ", errRes)
	if res == nil {
		logrus.Error("Config : Cannot start Program, Failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() (*ProgramConfig, error) {
	var errorLoad error
	var res = new(ProgramConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid Server Port Value, ", err.Error())
			permit = false
		}
		res.Server = port
	} else {
		permit = false
		errorLoad = errors.New("SERVER PORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid DB Port Value, ", err.Error())
			permit = false
		}
		res.DBPort = port
	} else {
		permit = false
		errorLoad = errors.New("DATABASE PORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
		errorLoad = errors.New("DATABASE HOST UNDEFINED")
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
		errorLoad = errors.New("DATABASE USER UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	} else {
		permit = false
		errorLoad = errors.New("DATABASE PASSWORD UNDEFINED")
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
		errorLoad = errors.New("DATABASE NAME UNDEFINED")
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
		errorLoad = errors.New("SECRET UNDEFINED")
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	} else {
		permit = false
		errorLoad = errors.New("REF SECRET UNDEFINED")
	}

	if !permit {
		return nil, errorLoad
	}

	return res, nil
}
