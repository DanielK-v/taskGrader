package models

import (
	"database/sql"

	"github.com/DanielK_v/taskGrader/services/database"
)

type Task struct {
	Id     uint64
	Name   string `json:"name" binding:"required"`
	Rating uint64 `json:"rating" binding:"required"`
}


func New(id uint64, name string, rating uint64) *Task {
	return &Task{
		Id:     id,
		Name:   name,
		Rating: rating,
	}
}

func (t Task) AddTask(task Task) (*Task, error) {
	connection, errCon := database.Connect()

	if errCon != nil {
		return nil, errCon
	}

	_, err := connection.Exec("INSERT INTO `tasks` (`name`, `rating`) VALUES (?, ?)", task.Name, task.Rating)

	if err != nil {
		return nil, err
	}

	defer connection.Close()

	return &task, nil
}

func GetAllTasks() ([]Task, error) {
	connection, errCon := database.Connect()

	if errCon != nil {
		panic(errCon)
	}

	rows, err := connection.Query("SELECT * FROM `tasks`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer connection.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Name, &task.Rating)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(id uint64) (*Task, error) {
	connection, errCon := database.Connect()

	if errCon != nil {
		panic(errCon)
	}

	row := connection.QueryRow("SELECT * FROM `tasks` WHERE id = ?", id)
	defer connection.Close()

	var task Task
	err := row.Scan(&task.Id, &task.Name, &task.Rating)

	if err == sql.ErrNoRows {
		return nil, err
	}

	return &task, nil
}

func DeleteTask(id uint64) error {
	connection, errCon := database.Connect()

	if errCon != nil {
		panic(errCon)
	}

	_, err := connection.Exec("DELETE FROM `tasks` WHERE id = ?", id)
	defer connection.Close()

	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(task *Task) error {
	connection, errCon := database.Connect()

	if errCon != nil {
		panic(errCon)
	}

	_, err := connection.Exec("UPDATE `tasks` SET name = ?, rating = ? WHERE id = ?", task.Name, task.Rating, task.Id)
	defer connection.Close()

	if err != nil {
		return err
	}

	return nil
}
