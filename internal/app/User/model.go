package User

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Service interface {
	Add(user *User) (uint64, error)
	Read(userId *uint64) (*User, error)
	Update(user *User, id *uint64) error
	Delete(userId *uint64) error
}

func (User) TableName() string {
	return "user"
}
