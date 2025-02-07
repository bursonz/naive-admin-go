package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"naive-admin-go/model"
	"os"
	"time"
)

var Dao *gorm.DB

func Init() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	)
	openDb, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("db connection error is %s", err.Error())
	}
	dbCon, err := openDb.DB()
	if err != nil {
		log.Fatalf("openDb.DB error is  %s", err.Error())
	}
	dbCon.SetMaxIdleConns(3)
	dbCon.SetMaxOpenConns(10)
	dbCon.SetConnMaxLifetime(time.Hour)
	Dao = openDb
}

func AutoMigrate() error {
	if err := Dao.AutoMigrate(
		&model.Permission{},
		&model.Profile{},
		&model.Role{},
		&model.RolePermissionsPermission{},
		&model.User{},
		&model.UserRolesRole{},
		&model.SysLog{},
		&model.Station{},
		&model.Lock{},
		&model.Order{},
		&model.OrderApproval{},
		&model.OrderStep{},
	); err != nil {
		return err
	}
	return nil
}
