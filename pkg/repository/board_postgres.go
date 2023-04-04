package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	money "github.com/vovah1a/my-money"
	"strings"
)

type BoardPostgres struct {
	db *sqlx.DB
}

func NewBoardPostgres(db *sqlx.DB) *BoardPostgres {
	return &BoardPostgres{db: db}
}

func (r *BoardPostgres) Create(userId int, board money.Board) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBoardQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", boardTable)
	row := tx.QueryRow(createBoardQuery, board.Title, board.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserBoardQuery := fmt.Sprintf("INSERT INTO %s (user_id, board_id) VALUES ($1, $2)", usersBoardTable)

	_, err = tx.Exec(createUserBoardQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BoardPostgres) GetAll(userId int) ([]money.Board, error) {
	var boards []money.Board

	query := fmt.Sprintf("SELECT b.id, b.title, b.description FROM %s b INNER JOIN %s ub on b.id = ub.board_id WHERE ub.user_id = $1",
		boardTable, usersBoardTable)

	err := r.db.Select(&boards, query, userId)
	return boards, err
}

func (r *BoardPostgres) GetById(userId, boardId int) (money.Board, error) {
	var board money.Board

	query := fmt.Sprintf(`SELECT b.id, b.title, b.description FROM %s b
								INNER JOIN %s ub on b.id = ub.board_id WHERE ub.user_id = $1 AND ub.board_id = $2`,
		boardTable, usersBoardTable)

	err := r.db.Get(&board, query, userId, boardId)
	return board, err
}

func (r *BoardPostgres) Delete(userId, boardId int) error {
	query := fmt.Sprintf("DELETE FROM %s b USING %s ub WHERE b.id=ub.board_id AND ub.user_id=$1 AND ub.board_id=$2",
		boardTable, usersBoardTable)
	_, err := r.db.Exec(query, userId, boardId)

	return err
}

func (r *BoardPostgres) Update(userId, boardId int, input money.UpdateBoardInput) error {
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

	query := fmt.Sprintf("UPDATE %s b SET %s FROM %s ub WHERE b.id = ub.board_id AND ub.board_id=$%d AND ub.user_id=$%d",
		boardTable, setQuery, usersBoardTable, argId, argId+1)
	args = append(args, boardId, userId)

	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
