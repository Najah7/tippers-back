package db

import "tippers-back/db/table"

func (d *DB) GetStaffByRestaurantID(id int) (*[]table.User, error) {
	var staffs []table.User
	if err := d.Conn.Where("is_employed = ? AND restaurant_id = ?", 1, id).Find(&staffs).Error; err != nil {
		return nil, err
	}
	return &staffs, nil
}
