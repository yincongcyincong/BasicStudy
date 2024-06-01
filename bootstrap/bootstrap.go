package bootstrap

import (
	"github.com/BurntSushi/toml"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"xorm.io/xorm"
)

type AppConfig struct {
	AppName    string `toml:"app_name"`
	HTTPListen string `toml:"http_listen"`

	DBFile   string `toml:"db_file"`
	DataFile string `toml:"data_file"`
}

var (
	AppConfigInstance = new(AppConfig)
	DB                *xorm.Engine
)

// InitConf init conf
func InitConf() {
	// init conf
	_, err := toml.DecodeFile("./conf/app.toml", AppConfigInstance)
	if err != nil {
		panic(err)
	}
}

// InitDB init database
func InitDB() {
	var err error
	dbExist := util.Exists(AppConfigInstance.DataFile)
	DB, err = xorm.NewEngine("sqlite3", AppConfigInstance.DataFile)
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(1)
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	if !dbExist {
		createDB()
	}
}

// createDB create database
func createDB() {
	data, err := util.ReadFile(AppConfigInstance.DBFile)
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec(data)
	if err != nil {
		panic(err)
	}
}
