package users

type UserRepository interface {
	GetAll() []User
	GetByID(di int) *User
}

type InMemoryUserRepository struct {
	users []User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []User{
			{ID: 1, Name: "Juan Perez", Email: "foo@foo.com"},
			{ID: 2, Name: "Ma Lo", Email: "bar@bar.com"},
		},
	}
}

func (r *InMemoryUserRepository) GetAll() []User {
	return r.users
}

func (r *InMemoryUserRepository) GetByID(id int) *User {
	for _, u := range r.users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}
