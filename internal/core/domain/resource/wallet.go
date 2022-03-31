package resource

import "time"

type Wallet struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Balance   string    `gorm:"type:varchar(32)" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
