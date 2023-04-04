package service

import (
	money "github.com/vovah1a/my-money"
	"github.com/vovah1a/my-money/pkg/repository"
)

type BoardService struct {
	repo repository.Board
}

func NewBoardService(repo repository.Board) *BoardService {
	return &BoardService{repo: repo}
}

func (s *BoardService) Create(userId int, board money.Board) (int, error) {
	return s.repo.Create(userId, board)
}

func (s *BoardService) GetAll(userId int) ([]money.Board, error) {
	return s.repo.GetAll(userId)
}

func (s *BoardService) GetById(userId, boardId int) (money.Board, error) {
	return s.repo.GetById(userId, boardId)
}

func (s *BoardService) Delete(userId, boardId int) error {
	return s.repo.Delete(userId, boardId)
}

func (s *BoardService) Update(userId, boardId int, input money.UpdateBoardInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, boardId, input)
}
