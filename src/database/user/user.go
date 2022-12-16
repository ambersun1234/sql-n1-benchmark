package user

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
}

func (u *User) TableName() string {
	return "user"
}
