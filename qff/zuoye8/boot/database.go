package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/app/global"
	"time"
)

func MysqlDBSetup() {
	config := global.Config.Database.Mysql

	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		global.Logger.Fatal("initialize database failed", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(global.Config.Database.Mysql.GetConnMaxIDleTime())
	sqlDB.SetConnMaxIdleTime(global.Config.Database.Mysql.GetConnMaxIDleTime())
	sqlDB.SetMaxIdleConns(global.Config.Database.Mysql.MaxOpenConns)
	sqlDB.SetMaxOpenConns(global.Config.Database.Mysql.MaxIdleConns)
	err = sqlDB.Ping()
	if err != nil {
		global.Logger.Fatal("connected failed", zap.Error(err))
	}
	global.MysqlDB = db

	global.Logger.Info("initialize database success")
}

func RedisSetup() {
	config := global.Config.Database.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s %s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("connect to redis failed", zap.Error(err))
	}
	global.Rdb = rdb

	global.Logger.Info("initialize redis success")
}
