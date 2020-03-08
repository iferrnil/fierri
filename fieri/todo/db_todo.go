package todo

import (
	"context"
	"database/sql"
	"log"
)

type DbTodo struct {
	GidLen int
	ctx    context.Context
	db     *sql.DB
}

func (dt *DbTodo) Add(what string) (res *ToDoItem, err error) {
	res = &ToDoItem{
		ToDo: what,
		Gid:  RandGid(dt.GidLen),
	}
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	result, err := db.Exec("insert into task (gid, todo, created) values($1, $2, current_timestamp)", res.Gid, res.ToDo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, nil
	}
	return res, nil
}

func NewDbTodo(gidLen int, ctx context.Context, db *sql.DB) *DbTodo {
	return &DbTodo{
		GidLen: gidLen,
		ctx:    ctx,
		db:     db,
	}
}

func (dt *DbTodo) init() {
	dt.Add("Dodać obsługę dodawania przez API")
	dt.Add("Dodać trwałe składowanie")
}

// dużo powtórek, ale to żeby zrozumieć
// poźniej do przemyślania jak to zrobić z klasą
// w szczególności transkacja - gdy nic nie zmieniamy ?

func (dt *DbTodo) Update(newValue *ToDoItem) (*ToDoItem, error) {
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	result, err := db.Exec("update task set todo = $1 where gid = $2", newValue.ToDo, newValue.Gid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, nil
	}
	return newValue, nil
}

func (dt *DbTodo) Remove(gid string) (*ToDoItem, error) {
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	removed, err := dt.findByGid(gid, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	result, err := db.Exec("delete from task where gid = $1", gid)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	if rowsAffected == 0 {
		return nil, nil
	}
	return removed, nil
}

// pobiera element - nie zarzadza transkacja
func (dt *DbTodo) findByGid(gid string, tx *sql.Tx) (result *ToDoItem, err error) {
	rows, err := tx.Query("select gid, todo from task where gid = $1", gid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		result = &ToDoItem{}
		if err := rows.Scan(&result.Gid, &result.ToDo); err != nil {
			return nil, err
		}
		break
	}
	if rows.Next() {
		log.Fatalf("To many rows returned for gid %s", gid)
	}
	return result, nil
}

func (dt *DbTodo) FindByGid(gid string) (result *ToDoItem, err error) {
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	result, err = dt.findByGid(gid, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return result, nil
}

func (dt *DbTodo) List() ([]ToDoItem, error) {
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select gid, todo from task")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	result := make([]ToDoItem, 0)
	for rows.Next() {
		current := ToDoItem{}
		if err := rows.Scan(&current.Gid, &current.ToDo); err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, current)
	}
	tx.Commit()
	return result, nil
}

func (dt *DbTodo) Size() (int, error) {
	db := dt.db
	tx, err := db.BeginTx(dt.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return -1, err
	}
	rows, err := db.Query("select count(*) from task")
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	for rows.Next() {
		var result int = 0
		if err := rows.Scan(&result); err != nil {
			tx.Rollback()
			return -1, err
		}
		if err := tx.Commit(); err != nil {
			return -1, err
		}
		return result, nil
	}
	if err := tx.Rollback(); err != nil {
		return -1, err
	}
	return -1, nil
}
