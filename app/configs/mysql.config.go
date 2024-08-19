package configs

import (
	"fmt"
	"time"

	"github.com/phucpham-infinity/go-nextpress/app/context"
	user_model "github.com/phucpham-infinity/go-nextpress/app/module/user/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func InitMysql() {
	logger := context.AppContext().GetLogger()
	config := context.AppContext().GetConfig().MySql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Info("Connected to database! " + config.DbName)
	context.AppContext().SetMySqlConnection(db)

	SetPoolMysql()
	// Migrate()
}

func GenDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./app/models",                                                     // output path
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	db := context.AppContext().GetMySqlConnection()
	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
}

func SetPoolMysql() {
	logger := context.AppContext().GetLogger()
	config := context.AppContext().GetConfig()

	db := context.AppContext().GetMySqlConnection()
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Error setting pool database", zap.Error(err))
		panic(err)
	}
	sqlDB.SetMaxIdleConns(config.MySql.MaxIdeConns)
	sqlDB.SetMaxOpenConns(config.MySql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MySql.ConnMaxLifetime))
	logger.Info("Set pool database success!")
}

func Migrate() {
	logger := context.AppContext().GetLogger()
	db := context.AppContext().GetMySqlConnection()
	if err := db.AutoMigrate(
		&user_model.User{},
		&user_model.UserMeta{},
	); err != nil {
		logger.Error("Error migrating database", zap.Error(err))
		panic(err)
	}
	logger.Info("Migrate database success!")
}
