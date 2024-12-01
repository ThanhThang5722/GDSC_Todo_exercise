package models

import "github.com/ThanhThang5722/GDSC_Todo_exercise/pkg/database"

type TodoStatus string

const (
	Doing TodoStatus = "Doing"
	Done  TodoStatus = "Done"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      string
}

type NewTodo struct {
	Title       string
	Description string
}

func GetAllTodos() ([]Todo, error) {
	var Todos []Todo
	db := database.GetDbInstance()
	query := `
		SELECT id, Title, Description, Status
		FROM TODO
	`
	results, err := db.Query(query)
	if err != nil {
		return Todos, err
	}

	for results.Next() {
		var todo Todo
		err = results.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)
		if err != nil {
			return Todos, err
		}
		Todos = append(Todos, todo)
	}
	return Todos, nil
}

func DeleteTodo(id string) error {
	db := database.GetDbInstance()
	query := `
		DELETE FROM TODO
		WHERE ID = ?
	`
	_, err := db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) Update() error {
	db := database.GetDbInstance()
	query := `
		UPDATE TODO
		SET Title = ?, Description = ?, Status = ?
		WHERE ID = ?
	`
	_, err := db.Query(query, t.Title, t.Description, t.Status, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *NewTodo) SaveNewTodo() error {
	db := database.GetDbInstance()
	query := `
		INSERT INTO TODO (Title, Description)
		Values (?,?)
	`
	_, err := db.Query(query, t.Title, t.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetTodoByID(id string) (Todo, error) {
	var todo Todo
	db := database.GetDbInstance()
	query := `
		SELECT Title, Description, Status
		FROM TODO
		WHERE ID = ?
	`
	results, err := db.Query(query, id)
	if err != nil {
		return todo, err
	}

	for results.Next() {
		var todo Todo
		err = results.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)
		if err != nil {
			return todo, err
		}
	}
	return todo, nil
}
