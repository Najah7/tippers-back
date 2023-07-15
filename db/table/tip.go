package table

import (
	"gorm.io/gorm"
)

type Tip struct {
	gorm.Model
	SenderID      int `gorm:"foreignkey:ID"`
	ReceiverID    int `gorm:"foreignkey:ID"`
	Amount        int
	ThanksMessage string
}
