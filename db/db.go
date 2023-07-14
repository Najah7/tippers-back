package db

import (
	"log"
	"os"
	"time"
	"tippers-back/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	conn  *gorm.DB
	limit uint
}

func NewDB() (*DB, error) {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_DBNAME")

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	var limit uint = 100
	return &DB{
		conn:  db,
		limit: limit,
	}, nil
}

func (d *DB) CreateTable() error {
	if err := d.conn.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	if err := d.conn.AutoMigrate(&model.Restaurant{}); err != nil {
		return err
	}
	if err := d.conn.AutoMigrate(&model.Tip{}); err != nil {
		return err
	}
	if err := d.conn.AutoMigrate(&model.PaypayID{}); err != nil {
		return err
	}
	return nil
}
