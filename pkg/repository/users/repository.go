package users

import (
	"todo-list/pkg/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}
type TodoListRepository interface {
	Create(userId int, list models.Todolist) (int, error)
	GetAll(userId int) ([]models.Todolist, error)
	GetById(userId, listId int) (models.Todolist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}
type TodoItemRepository interface {
}
type Repository struct {
	*AuthPostgres
	*TodoListsPostgres
	//*TodoItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthPostgres:      NewAuthPostgres(db),
		TodoListsPostgres: NewTodoListsPostgres(db),
	}
}
