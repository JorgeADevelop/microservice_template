package fx

import (
	"fmt"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"www.marawa.com/microservice_service/internal/infra/config"
)

func NewDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error opening database: %w", err))
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("error getting sql.DB from gorm.DB: %w", err))
	}

	if err := sqlDB.Close(); err != nil {
		panic(fmt.Errorf("error closing database: %w", err))
	}
}

var DBModule = fx.Options(
	fx.Provide(NewDB),
	fx.Invoke(CloseDB),
)
