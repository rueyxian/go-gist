package database

import (
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Register the postgres database/sql driver.
)

type Config struct {
	User       string
	Password   string
	Host       string
	Name       string
	DisableTLS bool
}

// Open opens a database database connection.
func Open(c Config) (*sqlx.DB, error) {
	q := url.Values{}

	if c.DisableTLS {
		q.Set("sslmode", "disable")
	} else {
		q.Set("sslmode", "require")
	}

	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     c.Host,
		Path:     c.Name,
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}

func Wait(db *sqlx.DB, maxAttempts int, increment time.Duration) error {
	var pingError error
	for i := 0; i < maxAttempts; i++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(i) * increment)
	}
	return pingError
}
