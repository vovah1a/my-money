package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	money "github.com/vovah1a/my-money"
	"strings"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) Create(boardId int, category money.Category) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var categoryId int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", categoryTable)

	row := tx.QueryRow(createCategoryQuery, category.Title, category.Description)
	err = row.Scan(&categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createBoardCategoryQuery := fmt.Sprintf("INSERT INTO %s (board_id, category_id) values ($1, $2)", boardsCategoryTable)
	_, err = tx.Exec(createBoardCategoryQuery, boardId, categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return categoryId, tx.Commit()
}

func (r *CategoryPostgres) GetAll(userId, boardId int) ([]money.Category, error) {
	var categories []money.Category
	qoery := fmt.Sprintf(`SELECT cat.id, cat.title, cat.description FROM %s cat INNER JOIN %s bc on bc.category_id = cat.id
    							INNER JOIN %s ub on ub.board_id = bc.board_id WHERE bc.board_id=$1 AND ub.user_id=$2`,
		categoryTable, boardsCategoryTable, usersBoardTable)

	if err := r.db.Select(&categories, qoery, boardId, userId); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryPostgres) GetById(userId, categoryId int) (money.Category, error) {
	var category money.Category

	query := fmt.Sprintf(`SELECT cat.id, cat.title, cat.description FROM %s cat
								INNER JOIN %s bc on cat.id = bc.category_id
								INNER JOIN %s ub on bc.board_id = ub.board_id WHERE cat.id = $1 AND ub.user_id = $2`,
		categoryTable, boardsCategoryTable, usersBoardTable)

	err := r.db.Get(&category, query, categoryId, userId)
	return category, err
}

func (r *CategoryPostgres) Delete(userId, categoryId int) error {
	query := fmt.Sprintf(`DELETE FROM %s cat USING %s bc, %s ub
       							WHERE cat.id=bc.category_id AND bc.board_id=ub.board_id AND ub.user_id=$1 AND cat.id=$2`,
		categoryTable, boardsCategoryTable, usersBoardTable)
	_, err := r.db.Exec(query, userId, categoryId)

	return err
}

func (r *CategoryPostgres) Update(userId, categoryId int, input money.UpdateCategoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s cat SET %s FROM %s bc, %s ub
                     			WHERE cat.id = bc.category_id AND bc.board_id=ub.board_id AND cat.id=$%d AND ub.user_id=$%d`,
		categoryTable, setQuery, boardsCategoryTable, usersBoardTable, argId, argId+1)
	args = append(args, categoryId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
