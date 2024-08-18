package configs

import (
	"fmt"
	"time"

	"github.com/phucpham-infinity/go-nextpress/app/global"
	"github.com/phucpham-infinity/go-nextpress/app/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	config := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		global.Logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Connected to database! " + config.DbName)
	global.DB = db

	SetPoolMysql()
	Migrate()
}

func SetPoolMysql() {
	sqlDB, err := global.DB.DB()
	if err != nil {
		global.Logger.Error("Error setting pool database", zap.Error(err))
		panic(err)
	}
	sqlDB.SetMaxIdleConns(global.Config.MySql.MaxIdeConns)
	sqlDB.SetMaxOpenConns(global.Config.MySql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.MySql.ConnMaxLifetime))
	global.Logger.Info("Set pool database success!")
}

func CloseMysql() {
	db, err := global.DB.DB()
	if err != nil {
		global.Logger.Error("Error closing database", zap.Error(err))
		panic(err)
	}
	db.Close()
	global.Logger.Info("Database connection closed")
}

func Migrate() {
	if err := global.DB.AutoMigrate(
		&po.User{},
		&po.UserMeta{},
	); err != nil {
		global.Logger.Error("Error migrating database", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Migrate database success!")
}
