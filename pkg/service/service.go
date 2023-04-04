package service

import (
	money "github.com/vovah1a/my-money"
	"github.com/vovah1a/my-money/pkg/repository"
)

type Authorization interface {
	CreateUser(user money.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Board interface {
	Create(userId int, list money.Board) (int, error)
	GetAll(userId int) ([]money.Board, error)
	GetById(userId, boardId int) (money.Board, error)
	Delete(userId, boardId int) error
	Update(userId, boardId int, input money.UpdateBoardInput) error
}

type Category interface {
	Create(userId, boardId int, category money.Category) (int, error)
	GetAll(userId, boardId int) ([]money.Category, error)
	GetById(userId, categoryId int) (money.Category, error)
	Delete(userId, categoryId int) error
	Update(userId, categoryId int, input money.UpdateCategoryInput) error
}

type Service struct {
	Authorization
	Board
	Category
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Board:         NewBoardService(repos.Board),
		Category:      NewCategoryService(repos.Category, repos.Board),
	}
}
