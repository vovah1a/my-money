package service

import (
	money "github.com/vovah1a/my-money"
	"github.com/vovah1a/my-money/pkg/repository"
)

type CategoryService struct {
	repo      repository.Category
	boardRepo repository.Board
}

func NewCategoryService(repo repository.Category, boardRepo repository.Board) *CategoryService {
	return &CategoryService{repo: repo, boardRepo: boardRepo}
}

func (s *CategoryService) Create(userId, boardId int, category money.Category) (int, error) {
	_, err := s.boardRepo.GetById(userId, boardId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(boardId, category)
}

func (s *CategoryService) GetAll(userId, boardId int) ([]money.Category, error) {
	return s.repo.GetAll(userId, boardId)
}

func (s *CategoryService) GetById(userId, categoryId int) (money.Category, error) {
	return s.repo.GetById(userId, categoryId)
}

func (s *CategoryService) Delete(userId, categoryId int) error {
	return s.repo.Delete(userId, categoryId)
}

func (s *CategoryService) Update(userId, categoryId int, input money.UpdateCategoryInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, categoryId, input)
}
