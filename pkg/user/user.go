package user

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
