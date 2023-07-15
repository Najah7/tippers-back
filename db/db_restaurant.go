package db

import "tippers-back/db/table"

func (d *DB) GetRestaurants() (*[]table.Restaurant, error) {
	var restaurants *[]table.Restaurant
	if err := d.Conn.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}
