package db

import (
	"tippers-back/db/table"
)

func (d *DB) GetTipsBySenderID(senderID int) (*[]table.Tip, error) {
	var tips *[]table.Tip
	if err := d.Conn.Where("sender_id = ?", senderID).Find(&tips).Error; err != nil {
		return nil, err
	}
	return tips, nil
}

func (d *DB) SendTip(tip *table.Tip) (*table.Tip, error) {

	if err := d.Conn.Create(&tip).Error; err != nil {
		return nil, err
	}
	return tip, nil
}

func (d *DB) GetTipAmountBySenderID(senderID int) (*int, error) {
	var result int
	row := d.Conn.Model(&table.Tip{}).Select("sum(amount)").Where("sender_id = ?", senderID).Row()
	if err := row.Scan(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
