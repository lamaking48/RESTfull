package sqlitelocal

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {

	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s,%w", op, err)
	}

	smtm, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	_, err = smtm.Exec()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "Storage.saqlite.SaveUrL"

	stmt, err := s.db.Prepare("INSERT INTO url(url,alias) VALUES (?,?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil {
		// Адаптируй это под sqlite а не sqlite3
		if sqliteErr, ok := err.(*sqlite.Error); ok && sqliteErr.Code() == sqlite.SQLITE_CONSTRAINT_UNIQUE {
			return 0, fmt.Errorf("%s: %w работает уникальность", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}
