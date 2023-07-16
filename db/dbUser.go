package db

import (
	"tippers-back/db/table"

	"gorm.io/gorm"
)

func (d *DB) GetUsers() (*[]table.User, error) {
	var users *[]table.User
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

func (d *DB) DeleteUserByID(id int) error {
	if err := d.Conn.Delete(&table.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateUserRestaurantIDByID(id, restaurantID int) error {
	if err := d.Conn.Model(&table.User{}).Where("id = ?", id).Update("restaurant_id", restaurantID).Error; err != nil {
		return err
	}
	return nil

}

func (d *DB) UpdateUserMoneyByID(senderID, receiverID, money int) error {
	if err := d.Conn.Model(&table.User{}).Where("id = ?", senderID).Update("money", gorm.Expr("money - ?", money)).Error; err != nil {
		return err
	}

	if err := d.Conn.Model(&table.User{}).Where("id = ?", receiverID).Update("money", gorm.Expr("money + ?", money)).Error; err != nil {
		return err
	}
	return nil

}

func (d *DB) GetEmployedUsers() (*[]table.User, error) {
	var users []table.User
	if err := d.Conn.Where("is_employed = ?", 1).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
