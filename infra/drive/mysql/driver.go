package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
)

type MysqlDriver struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
	client       *gorm.DB
}

func (*MysqlDriver) Configure(ConfigPath string, ConfigName string) MysqlDriver {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("default.host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}

	response := MysqlDriver{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
		"",
		config.GetBool("default.sql_log"),
	}

	return response
}

func (c *MysqlDriver) NewDriver() error {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		return err

	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

	if c.print_log {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}
	c.client = db
	return nil
}
