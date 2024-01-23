package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)


func (app *App) routes() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/register", app.register).Methods(http.MethodPost).Headers("Accept", "application/json")

	// 路由分组
	book := router.PathPrefix("/books").Subrouter()
	{
		// bug：`/books` 和 `/books/`，两者是否有区别
		book.HandleFunc("/{title}", app.GetBook).Methods("GET")
	}

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := jsonResponse {
			Error:   false,
			Message: "请求页面未找到:( 如有疑惑，请联系我们。",
		}

		// 格式化输出
		// 前缀，"" 无
		// 缩进，"\t" 
		out, _ := json.MarshalIndent(payload, "", "\t")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(out)
	})

	router.Use(cors)
	router.Use(logging)

	return router  
}





/*

// 写法 1
func (app *Config) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}



*/



// 写法 3
func (app *App) Ping(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse {
		Error:   false,
		Message: "Pong",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}


func (app *App) notFound(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse {
		Error: false,
		Message: "请求页面未找到:( 如有疑惑，请联系我们。",
	}

	_ = app.writeJSON(w, http.StatusNotFound, payload)
}


