package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"sync"
    "flag"
)

var (
	dbConfigFile = "db-config.json"
	uiConfigFile = "ui-config.json"
)

type DBConfig struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	DbName       string `json:"dbName"`
	DbDriverName string `json:"dbDriver"`
}

type UiConfig map[string]interface{}

var config *DBConfig
var uiConfig *UiConfig

var dbCfgOnce sync.Once
var uiCfgOnce sync.Once

func GetDBConfig() *DBConfig {
	dbCfgOnce.Do(func() {
		config = readConfig(DBConfig{}).(*DBConfig)
	})
	return config
}

func GetUIConfig() *UiConfig {
	uiCfgOnce.Do(func() {
		uiConfig = readConfig(UiConfig{}).(*UiConfig)
	})
	return uiConfig
}

func readConfig(cfg interface{}) interface{} {
    readFromFlags()

    //
	switch reflect.ValueOf(cfg).Interface().(type) {
	case DBConfig:
		db := new(DBConfig)
		f, err := ioutil.ReadFile(dbConfigFile)
		err = json.Unmarshal(f, db)
		if err != nil {
			log.Fatalf("read DB config error: %s, %#v", err, db)
		}
		return db
	case UiConfig:
		ui := new(UiConfig)
		f, err := ioutil.ReadFile(uiConfigFile)
		err = json.Unmarshal(f, ui)
		if err != nil {
			log.Fatalf("read UI config error: %s, %#v", err, ui)
		}
		return ui
	default:
		log.Fatalf("can't specify config file type: %s", cfg)
	}
	return nil
}

func readFromFlags()  {
    dbCfg := flag.String("db-config", "db-config.json", "Path to DB config")
    uiCfg := flag.String("ui-config", "ui-config.json", "Path to DB config")
    flag.Parse()

    if dbCfg != nil {
        dbConfigFile = *dbCfg
    }

    if uiCfg != nil {
        uiConfigFile = *uiCfg
    }
}
