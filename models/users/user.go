package models

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

var users []User

func New(id int, username, password, email string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

func AddUser(user User) {
	// add users
}

func GetAllUsers() {
	// Fetch users from DB
}
