package Model

type UsersLists struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	UserID uint `gorm:"not null"`
	ListID uint `gorm:"not null"`
}
