package db

import "tippers-back/db/table"

func (d *DB) GetTipsBySenderID(senderID int) (*[]table.Tip, error) {
	var tips *[]table.Tip
	if err := d.Conn.Where("sender_id = ?", senderID).Find(&tips).Error; err != nil {
		return nil, err
	}
	return tips, nil
}