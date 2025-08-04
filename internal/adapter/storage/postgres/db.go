package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	User string
	Pwd  string
	Host string
	Port uint16
	Name string
}

type Database interface {
	GetInstance() *gorm.DB
	Connect()
	Disconnect()
}

type database struct {
	db      *gorm.DB
	context context.Context
	config  DbConfig
}

func NewDatabase(ctx context.Context, config DbConfig) Database {
	db := database{
		context: ctx,
		config:  config,
	}
	return &db
}

func (d *database) Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.config.Host,
		d.config.User,
		d.config.Pwd,
		d.config.Name,
		d.config.Port,
	)

	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("connected to database!")

	err = d.db.AutoMigrate(&model.StudentModel{}, &model.TeacherModel{}, &model.ClassModel{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

}

func (d *database) Disconnect() {
	sqlDB, err := d.db.DB()
	if err != nil {
		log.Printf("error getting database instance: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("error disconnecting database: %v", err)
	} else {
		log.Println("disconnected database")
	}
}

func (d *database) GetInstance() *gorm.DB {
	return d.db
}
