package database

import (
	"database/sql"
	"log"
	"time"

	"server/model"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "questions.db")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		type TEXT NOT NULL,
		options TEXT NOT NULL,
		answer TEXT NOT NULL,
		difficulty TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		deleted_at TEXT
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("建表失败:", err)
	}
}

func GetAllQuestions() ([]model.Question, error) {
	rows, err := DB.Query("SELECT id, title, type, options, answer, difficulty, created_at, updated_at, deleted_at FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []model.Question
	for rows.Next() {
		var q model.Question
		var deleted sql.NullString
		err := rows.Scan(&q.ID, &q.Title, &q.Type, &q.Options, &q.Answer, &q.Difficulty, &q.CreatedAt, &q.UpdatedAt, &deleted)
		if err != nil {
			return nil, err
		}
		if deleted.Valid {
			q.DeletedAt = &deleted.String
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func InsertQuestion(q model.Question) error {
	now := time.Now().Format(time.RFC3339)
	_, err := DB.Exec(
		"INSERT INTO questions (title, type, options, answer, difficulty, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		q.Title, q.Type, q.Options, q.Answer, q.Difficulty, now, now,
	)
	return err
}

func InsertQuestionReturnID(q model.Question) (int64, error) {
	now := time.Now().Format(time.RFC3339)
	result, err := DB.Exec(
		"INSERT INTO questions (title, type, options, answer, difficulty, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		q.Title, q.Type, q.Options, q.Answer, q.Difficulty, now, now,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateQuestion(q model.Question) error {
	now := time.Now().Format(time.RFC3339)
	_, err := DB.Exec(
		"UPDATE questions SET title = ?, type = ?, options = ?, answer = ?, difficulty = ?, updated_at = ? WHERE id = ?",
		q.Title, q.Type, q.Options, q.Answer, q.Difficulty, now, q.ID,
	)
	return err
}

func DeleteQuestions(ids []int) error {
	stmt, err := DB.Prepare("DELETE FROM questions WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, id := range ids {
		if _, err := stmt.Exec(id); err != nil {
			return err
		}
	}
	return nil
}

func QueryQuestions(keyword string, questionType string, page, pageSize int) ([]model.Question, int, error) {
	offset := (page - 1) * pageSize

	// 构造条件
	where := "WHERE 1=1"
	args := []interface{}{}

	if keyword != "" {
		where += " AND title LIKE ?"
		args = append(args, "%"+keyword+"%")
	}
	if questionType != "" {
		where += " AND type = ?"
		args = append(args, questionType)
	}

	// 查询总数
	countSQL := "SELECT COUNT(*) FROM questions " + where
	var total int
	if err := DB.QueryRow(countSQL, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// 查询数据
	sqlStr := `
		SELECT id, title, type, options, answer, difficulty, created_at, updated_at, deleted_at
		FROM questions
		` + where + `
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`
	args = append(args, pageSize, offset)

	rows, err := DB.Query(sqlStr, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var questions []model.Question
	for rows.Next() {
		var q model.Question
		var deleted sql.NullString
		err := rows.Scan(&q.ID, &q.Title, &q.Type, &q.Options, &q.Answer, &q.Difficulty, &q.CreatedAt, &q.UpdatedAt, &deleted)
		if err != nil {
			return nil, 0, err
		}
		if deleted.Valid {
			q.DeletedAt = &deleted.String
		}
		questions = append(questions, q)
	}
	return questions, total, nil
}
