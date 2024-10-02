package main

// TODO : 65 - 70
// TODO : 89 - better model config - the gist
//

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type app struct {
	logger *slog.Logger
}

func main() {

	port := flag.String("p", ":8080", "serve port")
	db_url := flag.String("db", "user:password@/dbname", "db url : user:password@/dbname")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*db_url)
	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()
	app := &app{logger: logger}
	// TODO: from the videos go about &

	logger.Info("> server running on", "port", *port)

	err = http.ListenAndServe(*port, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(db_url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", db_url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil

}
