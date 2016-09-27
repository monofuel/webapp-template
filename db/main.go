package db

//Database wrapper
type Database interface {
	//List all Todos
	ListTodo() ([]string, error)

	//Add a Todo
	AddTodo(text string) (int64, error)

	//get a Todo
	GetTodo(id int64) (string, error)

	//Update a Todo
	UpdateTodo(id int64, text string) error

	//Delete a Todo
	DeleteTodo(id int64) error
	//Disconnect from the database
	Close()
}
