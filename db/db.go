package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"tippers-back/db/table"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn  *gorm.DB
	Limit uint
}

func NewDB() (*DB, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_TZ := os.Getenv("DB_TZ")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%s", DB_USER, DB_PASS, DB_PORT, DB_NAME, DB_TZ)

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
		Conn:  db,
		Limit: limit,
	}, nil
}

func (d *DB) GetUsers(users *[]table.User) (*[]table.User, error) {
	if err := d.Conn.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DB) GetUserByID(id int) (*table.User, error) {
	var user table.User
	if err := d.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DB) CreateTable() error {
	if err := d.Conn.AutoMigrate(&table.User{}); err != nil {
		return err
	}
	if err := d.Conn.AutoMigrate(&table.Restaurant{}); err != nil {
		return err
	}
	if err := d.Conn.AutoMigrate(&table.Tip{}); err != nil {
		return err
	}
	if err := d.Conn.AutoMigrate(&table.PaypayID{}); err != nil {
		return err
	}
	return nil
}

func (d *DB) GetUserByMail(mail string) (*table.User, error) {
	var user table.User
	if err := d.Conn.Where("email = ?", mail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DB) RegisterUser(user *table.User) (*table.User, error) {
	if err := d.Conn.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DB) UpdateUser(user *table.User) (*table.User, error) {
	if err := d.Conn.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
