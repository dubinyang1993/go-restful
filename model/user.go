package model

type User struct {
	ID    int
	Name  string
	Phone string
}

func (user *User) Create() error {
	return nil
}
