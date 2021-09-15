package main

import (
	"fmt"
	"go-gist/topic/design/cmd-docker-run-db/database"
	"go-gist/topic/design/cmd-docker-run-db/docker"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// ================================================================================
// main
func main() {

	if err := run(); err != nil {
		log.Fatalf(err.Error())
	}

}

// ================================================================================
// run
func run() error {

	log := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	db, teardown, err := NewUnit(log)
	if err != nil {
		return err
	}
	defer teardown()

	log.Println("access db ....", db)
	time.Sleep(3 * time.Second)

	return nil
}

// ================================================================================
// NewUnit
func NewUnit(log *log.Logger) (*sqlx.DB, func(), error) {

	c, err := docker.NewContainer()
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating container")
	}
	log.Printf("container id: %s\n", c.ID())

	if err := c.StartContainer(); err != nil {
		return nil, nil, errors.Wrap(err, "starting container")
	}
	log.Printf("db host: %s\n", c.Host())

	// ==============================
	db, err := database.Open(database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       c.Host(),
		Name:       "postgres",
		DisableTLS: true,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "opening database connection")
	}

	// ==============================
	if err := database.Wait(db, 20, 50*time.Millisecond); err != nil {
		b, err := c.DumpContainerLogs()
		if err != nil {
			return nil, nil, err
		}
		fmt.Println("========================================")
		log.Println(string(b))
		fmt.Println("========================================")

		defer c.RemoveContainer()
		defer c.StopContainer()
		return nil, nil, err
	}

	// ==============================
	teardown := func() {
		defer c.RemoveContainer()
		defer c.StopContainer()
		defer db.Close()
	}

	return db, teardown, nil
}
