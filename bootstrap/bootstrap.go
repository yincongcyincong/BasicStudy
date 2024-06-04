package bootstrap

import (
	"flag"
	"github.com/BurntSushi/toml"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"os"
	"xorm.io/xorm"
)

type AppConfig struct {
	AppName    string `toml:"app_name"`
	HTTPListen string `toml:"http_listen"`

	DBFile   string `toml:"db_file"`
	DataFile string `toml:"data_file"`

	ChatType string `toml:"chat_type"`

}

type GPT struct {
	SecretKey string `toml:"secret_key"`
}

type ErNie struct {
	Ak string `toml:"ak"`
	Sk string `toml:"sk"`
}

var (
	AppConfigInstance = new(AppConfig)
	GPTConfigInstance = new(GPT)
	ErNieConfigInstance = new(ErNie)
	DB                *xorm.Engine
)

// InitConf init conf
func InitConf() {
	// init conf
	_, err := toml.DecodeFile("./conf/app.toml", AppConfigInstance)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile("./conf/gpt.toml", GPTConfigInstance)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile("./conf/ernie.toml", ErNieConfigInstance)
	if err != nil {
		panic(err)
	}

	flag.Parse()

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
	DB.SetMaxOpenConns(10)
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	if !dbExist {
		createDB()
	}
}

func InitLog() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	file, err := os.OpenFile("./log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
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
