package entities

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserName  string `json:"userName"`
	LimitTask uint   `gorm:"not null" json:"limitTask"`
}

func (b *User) TableName() string {
	return "users"
}
