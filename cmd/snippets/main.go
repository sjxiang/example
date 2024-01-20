package main

import (
	"database/sql"
	"log"
	"flag"

	"github.com/sjxiang/example/models"
	_ "github.com/go-sql-driver/mysql"
)


type application struct {
	snippets       models.SnippetModelInterface
	users          models.UserModelInterface

	// lifetime
}

func main() {
	dsn := flag.String("dsn", "root:secret@(localhost:3306)/web?utf8mb4&parseTime=true&loc=Local", "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	app := &application{
		snippets: &models.SnippetModel{DB: db},
		users:    &models.UserModel{DB: db},
	}

	_  = app

	// 注册路由

}


func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}


/*

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			404
		} else {
			500
		}
		return
	}

 */
 

