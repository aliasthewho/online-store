package users

type User struct {
	ID    int
	Name  string
	Email string
}

func (u *User) GetInfo() string {
	return u.Name + " (" + u.Email + ")"
}
