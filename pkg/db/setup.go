package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"latentlab.cc/easyhousing/pkg/model"
	"os"
	"time"
)

func Setup() error {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
	}

	psql := postgres.New(cfg)

	gormCfg := gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(psql, &gormCfg)
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Home{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.SupportingDoc{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
