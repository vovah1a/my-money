package repository

import (
	"github.com/jmoiron/sqlx"
	money "github.com/vovah1a/my-money"
)

type Authorization interface {
	CreateUser(user money.User) (int, error)
	GetUser(username, password string) (money.User, error)
}

type Board interface {
	Create(userId int, board money.Board) (int, error)
	GetAll(userId int) ([]money.Board, error)
	GetById(userId, boardId int) (money.Board, error)
	Delete(userId, boardId int) error
	Update(userId, boardId int, input money.UpdateBoardInput) error
}

type Category interface {
	Create(boardId int, category money.Category) (int, error)
	GetAll(userId, boardId int) ([]money.Category, error)
	GetById(userId, categoryId int) (money.Category, error)
	Delete(userId, categoryId int) error
	Update(userId, categoryId int, input money.UpdateCategoryInput) error
}

type Repository struct {
	Authorization
	Board
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Board:         NewBoardPostgres(db),
		Category:      NewCategoryPostgres(db),
	}
}
