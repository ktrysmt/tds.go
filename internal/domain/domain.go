package domain

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}
