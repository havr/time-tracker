package stores

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/havr/time-tracker/internal/models"
)

type DatabaseSessionStore struct {
	db *sql.DB
}

type SessionStore interface {
	SaveSession(session *models.Session) error
	ListSessions(oldest time.Time) ([]models.Session, error)
}

func NewDatabaseSessionStore(db *sql.DB) *DatabaseSessionStore {
	return &DatabaseSessionStore{db: db}
}

func (s *DatabaseSessionStore) SaveSession(session *models.Session) error {
	_, err := s.db.Exec(`
			INSERT INTO 
				work_sessions (id, name, start_time, duration) 
			VALUES 
				($1, $2, $3, $4)
		`,
		session.ID,
		session.Name,
		session.StartTime,
		session.Duration,
	)
	if err != nil {
		return fmt.Errorf("save session: %s", err)
	}

	return nil
}

func (s *DatabaseSessionStore) ListSessions(limit time.Time) ([]models.Session, error) {
	rows, err := s.db.Query(`
		SELECT 
			id, name, start_time, EXTRACT(EPOCH FROM duration)
		FROM 
			work_sessions
		WHERE
			start_time > $1
		ORDER BY 
			start_time DESC
	`, limit)
	if err != nil {
		return nil, fmt.Errorf("query sessions: %s", err)
	}

	defer rows.Close()

	var sessions []models.Session

	for rows.Next() {
		var session models.Session
		if err := rows.Scan(&session.ID, &session.Name, &session.StartTime, &session.Duration); err != nil {
			return nil, fmt.Errorf("retreive session row: %s", err)
		}

		sessions = append(sessions, session)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close session rows: %s", err)
	}

	return sessions, nil
}
