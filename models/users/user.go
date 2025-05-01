package models

import "github.com/DanielK_v/taskGrader/services/database"

type User struct {
	Id       int    
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func New(id int, username, password, email string) *User {
	return &User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}

func AddUser(user User) (*User, error) {
	_, err := database.Db.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		 user.Username,
		 user.Email,
		 user.Password,
	)

	if err != nil {
        return nil, err
    }

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	rows, err := database.Db.Query("SELECT * FROM `users`")

	if err != nil {
		return nil, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
