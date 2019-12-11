package stores_test

import (
	"database/sql"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/havr/time-tracker/internal/models"
	"github.com/havr/time-tracker/internal/stores"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type TestFunc func(t *testing.T, store stores.SessionStore)

func TestStores(t *testing.T) {
	url := os.Getenv("TT_TEST_DB")
	if url == "" {
		t.Skip("no test database provided")
	}

	tests := map[string]TestFunc{
		"insertAndList":             testInsertAndList,
		"dontListOlderThanRequired": testDontListOlderThanRequired,
	}

	db, err := sql.Open("postgres", url)
	require.NoError(t, err)
	cleanDatabase(t, db) // clean any sample data left from migrations
	defer db.Close()

	store := stores.NewDatabaseSessionStore(db)

	for name, test := range tests {
		t.Run(name, func(subt *testing.T) {
			defer cleanDatabase(t, db)

			test(subt, store)
		})
	}
}

func cleanDatabase(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM work_sessions")
	require.NoError(t, err)
}

const (
	numSessionsToSpawn = 3
)

func testInsertAndList(t *testing.T, store stores.SessionStore) {
	var sessions []models.Session
	now := time.Now().In(time.Local)
	for i := 0; i < numSessionsToSpawn; i++ {
		session := models.Session{
			ID:        uuid.New(),
			Name:      strconv.Itoa(i),
			StartTime: now.Add(-time.Duration(i) * time.Second),
			Duration:  1,
		}

		err := store.SaveSession(&session)
		require.NoError(t, err)

		sessions = append(sessions, session)
	}

	result, err := store.ListSessions(time.Now().Add(-time.Hour))
	require.NoError(t, err)

	for i, session := range result {
		// db returns the session timezone, make it match the local one
		result[i].StartTime = session.StartTime.In(time.Local)
	}
	require.Equal(t, sessions, result)
}

func testDontListOlderThanRequired(t *testing.T, store stores.SessionStore) {
	limit := time.Now().Add(-time.Minute)
	err := store.SaveSession(&models.Session{
		ID:        uuid.New(),
		StartTime: limit.Add(-time.Hour),
		Duration:  1,
	})
	require.NoError(t, err)

	result, err := store.ListSessions(limit)
	require.NoError(t, err)
	require.Len(t, result, 0)
}
