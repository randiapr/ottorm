package ottorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/randiapr/ottorm/constant"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDbConnection(
	database, env, host, user, password, name, port, sslMode string,
	cfg ...gorm.Config) (*gorm.DB, error) {
	// init connection db with default postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, user, password, name, port, sslMode)
	if database == constant.GORM_MYSQL {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, name)
	}
	gormCfg := new(gorm.Config)
	if len(cfg) > 0 {
		gormCfg = &cfg[0]
	}
	// print sql statement when in development mode
	if env != constant.ENV_PROD {
		gormLog := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
		gormCfg.Logger = gormLog
	}
	// init gorm
	var (
		db  *gorm.DB
		err error
	)
	if database == constant.GORM_MYSQL {
		db, err = gorm.Open(mysql.Open(dsn), gormCfg)
	} else {
		db, err = gorm.Open(postgres.Open(dsn), gormCfg)
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
