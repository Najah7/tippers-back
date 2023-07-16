package db

import "tippers-back/db/table"

func (d *DB) GetRestaurants() (*[]table.Restaurant, error) {
	var restaurants *[]table.Restaurant
	if err := d.Conn.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (d *DB) GetRestaurantByID(id int) (*table.Restaurant, error) {
	var restaurant table.Restaurant
	if err := d.Conn.Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (d *DB) RegisterRestaurant(restaurant *table.Restaurant) (*table.Restaurant, error) {
	if err := d.Conn.Create(&restaurant).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (d *DB) UpdateRestaurantProfileImageURLIDByID(id int, pass string) error {
	if err := d.Conn.Model(&table.Restaurant{}).Where("id = ?", id).Update("profile_image_url", pass).Error; err != nil {
		return err
	}
	return nil
}
