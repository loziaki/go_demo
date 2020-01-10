package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"loziaki/go_demo/config"
	"os"
)

var DB *gorm.DB

func initDB() {
	Driver := config.Conf.Get("database.Driver").(string)
	Database := config.Conf.Get("database.Database").(string)
	Host := config.Conf.Get("database.Host").(string)
	User := config.Conf.Get("database.User").(string)
	Password := config.Conf.Get("database.Password").(string)
	Port := config.Conf.GetDefault("database.Port", 3306).(int64)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Database)
	fmt.Println(dataSourceName)
	db, err := gorm.Open(Driver, dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	DB = db
}

func init() {
	initDB()
}
