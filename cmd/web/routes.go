package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)


func (app *Config) routes() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/login", app.Login).Methods("POST")

	// 路由分组
	book := r.PathPrefix("/books").Subrouter()
	{
		// bug：`/books` 和 `/books/`，两者是否有区别
		book.HandleFunc("/{title}", app.GetBook).Methods("GET")
	}

	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	r.Use(cors)
	r.Use(logging)

	return r  
}



/*

// 写法 1
func (app *Config) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}



*/



// 写法 3
func (app *Config) Ping(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse {
		Error: false,
		Message: "Pong",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}


func (app *Config) notFound(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse {
		Error: false,
		Message: "请求页面未找到:( 如有疑惑，请联系我们。",
	}

	_ = app.writeJSON(w, http.StatusNotFound, payload)
}


