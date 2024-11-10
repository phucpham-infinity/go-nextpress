package configs

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/phucpham-infinity/go-nextpress/app/context"
	"go.uber.org/zap"
)

func InitMysql() {
	logger := context.AppContext().GetLogger()
	config := context.AppContext().GetConfig().Sql

	// dsn := "%s:%s@tcp(%s:%v)/%s"
	// dsn = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DbName)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Info("Connected to database! " + config.DbName)
	context.AppContext().SetMySqlConnection(db)

	SetPoolSql()
}

func SetPoolSql() {
	logger := context.AppContext().GetLogger()
	config := context.AppContext().GetConfig()

	db := context.AppContext().GetMySqlConnection()
	db.SetMaxIdleConns(config.Sql.MaxIdeConns)
	db.SetMaxOpenConns(config.Sql.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.Sql.ConnMaxLifetime))
	logger.Info("Set pool database success!")
}
