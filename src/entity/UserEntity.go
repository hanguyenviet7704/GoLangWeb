package entity

type User struct {
	Id       int    `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	RoleId   int    `json:"role_id"`
}
