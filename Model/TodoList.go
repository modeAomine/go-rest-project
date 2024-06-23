package Model

type TodoList struct {
	ID          uint   `gorm:"primary_key;auto_increment"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}
