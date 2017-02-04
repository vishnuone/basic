package main

import "encoding/json"
import "fmt"
import "os"
import "io/ioutil"

var configFilePath string // JSON file
var appConfig config

type confMySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type confRedis struct {
	Host string
}

type confDB struct {
	MySQLWriter confMySQL
	MySQLReader confMySQL
	RedisWriter confRedis
	RedisReader confRedis
}

type config struct {
	DB confDB
}

func init() {
	configFilePath = `./config.json` //@todo check if file exist in path. Else fetch path from ENV. Should use ENV for live.
	loadConfig()
}

func loadConfig() {
	appConfig = config{}
	contents := getFileContents(configFilePath)
	e := json.Unmarshal(contents, &appConfig)
	if e != nil {
		fmt.Println("Unable to read json from:" + configFilePath)
		os.Exit(1)
	}
}

func getFileContents(filePath string) (contents []byte) {
	contents, e := ioutil.ReadFile(filePath)
	if e != nil {
		fmt.Println("Unable to read file:" + filePath)
		os.Exit(1)
	}
	return
}
