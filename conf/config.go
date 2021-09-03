package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var shared *_Configuration

type _Configuration struct {
	Environment		string	`json:"environment"`
	AuthCred struct {
		Development struct {
			Username	string	`json:"username"`
			Password	string	`json:"password"`
		}	`json:"development"`
		Staging struct {
			Username	string	`json:"username"`
			Password	string	`json:"password"`
		}	`json:"staging"`
		Production struct {
			Username	string	`json:"username"`
			Password	string	`json:"password"`
		}	`json:"production"`
	}	`json:"auth_cred"`
}

func Init() {
	if shared != nil {
		return
	}

	fmt.Println("Config init")

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Base Path : " + basePath)

	config, err := ioutil.ReadFile(filepath.Join(basePath, "conf", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	shared = new(_Configuration)
	err = json.Unmarshal(config, &shared)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *shared
}
