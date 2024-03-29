package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sjxiang/example/models"
	"github.com/sjxiang/example/pkg/logger"
)

type Config struct {

}

type App struct {
	snippets  models.SnippetModelInterface
	users     models.UserModelInterface
	conf      Config
	// lifetime
}

var (
	dsn = flag.String("dsn", "root:secret@(localhost:3306)/web?utf8mb4&parseTime=true&loc=Local", "MySQL data source name")
)

type User struct {
	Name     string 
	Age      int
	Password string
}

func (u User) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "xxx"),
	)
} 

func main() {
	flag.Parse()

	logger.InitLogger()

	user := User{
		Name:     "sjxiang",
		Age:      28,
		Password: "123456",
	}

	slog.Info("starting api")
	slog.Info("creating user", "user", user.LogUser())

	db, err := openDB(*dsn)
	if err != nil {
		slog.Error("failed to connect database", "err", err, slog.String("package", "main"))
	}
	defer db.Close()


	app := &App{
		snippets: &models.SnippetModel{DB: db},
		users:    &models.UserModel{DB: db}, 
	}

	router := gin.Default()
	router.POST("/v1/user", app.Register)

	srv := &http.Server{
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("failed to start server", err, slog.String("package", "main"))
	}

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
 

