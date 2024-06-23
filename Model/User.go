package Model

type User struct {
	ID           uint   `gorm:"primary_key;auto_increment"`
	Name         string `gorm:"type:varchar(255);not null"`
	Username     string `gorm:"type:varchar(255);not null;unique"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
}
