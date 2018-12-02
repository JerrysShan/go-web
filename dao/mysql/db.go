package mysql

import (
	"go-web/utils"

	"github.com/go-xorm/xorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver
	logger "github.com/sirupsen/logrus"
)

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", utils.Config.GetString("mysql.url"))
	if err != nil {
		panic(err)
	}
	db.ShowSQL(true)
	logger.Info("mysql connect success")
}
