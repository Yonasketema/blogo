package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yonasketema/blogo/internal/models"
)

type app struct {
	logger        *slog.Logger
	blogs         *models.BlogModel
	templateCache map[string]*template.Template
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

	templateCache, err := templateCache()

	app := &app{
		logger:        logger,
		blogs:         &models.BlogModel{DB: db},
		templateCache: templateCache,
	}

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
