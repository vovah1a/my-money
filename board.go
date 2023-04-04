package money

import "errors"

type Board struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersBoard struct {
	Id      int
	UserId  int
	BoardId int
}

type Category struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type BoardsCategory struct {
	Id         int
	BoardId    int
	CategoryId int
}

type UpdateBoardInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateCategoryInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateCategoryInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i UpdateBoardInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
