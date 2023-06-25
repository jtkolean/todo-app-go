package store

import (
	"database/sql"
	"jtkolean/task/model"
	"log"
)

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (store *TaskStore) GetAll() ([]model.Task, error) {
	var tasks []model.Task

	log.Printf("scanned %v", len(tasks))

	rows, err := store.db.Query("select * from task")
	defer rows.Close()
	if err != nil {
		log.Printf("get all tasks failed from query, %v", err.Error())
		return tasks, err
	}

	var t model.Task
	for rows.Next() {
		if err := rows.Scan(&t.Id, &t.Title, &t.Completed, &t.CreateTs); err != nil {
			log.Printf("get all tasks failed from scan, %v", err.Error())
			break
		}
		tasks = append(tasks, t)
	}
	return tasks, err
}

func (store *TaskStore) Create(t model.Task) error {
	_, err := store.db.Exec("insert into task(title, completed, create_ts) values($1, $2, CURRENT_TIMESTAMP) returning id", t.Title, t.Completed)

	if err != nil {
		log.Printf("create task failed from insert, %v", err.Error())
	}

	return err
}

func (store *TaskStore) Update(id string, t model.Task) error {
	_, err := store.db.Exec("update task set title=$1, completed=$2 where id=$3", t.Title, t.Completed, id)

	if err != nil {
		log.Printf("update task failed from update id=%v, %v", id, err.Error())
	}

	return err
}

func (store *TaskStore) Delete(id string) error {
	_, err := store.db.Exec("delete from task where id = $1", id)

	if err != nil {
		log.Printf("delete task failed from delete id=%v, %v", id, err.Error())
	}

	return err
}
