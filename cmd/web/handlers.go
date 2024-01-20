package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// ***

	payload := jsonResponse{
		Error:   false,
		Message: "Hello world",
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}


func (app *Config) register(w http.ResponseWriter, r *http.Request) {
	// 表单
	name := r.FormValue("username")
	pwd  := r.FormValue("password")

	if len(name) == 0 || len(pwd) == 0 {
		app.errorJSON(w, errors.New("请求参数为空"), http.StatusBadRequest)
		return
	} 

	payload := jsonResponse{
		Error:   false,
		Message: "注册成功",
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}


func fetchVideos(w http.ResponseWriter, r *http.Request) {
	// e.g. `/videos?size=10&&number=1`

	// 参数路由
	_ = r.URL.Query().Get("size")    // 每页包含的数据条目
	_ = r.URL.Query().Get("number")  // 当前页码
}


func (app *Config) GetBook(w http.ResponseWriter, r *http.Request) {
	// e.g. `/books/the_golden_lotus`

	// 动态路由
	title := mux.Vars(r)["title"]
	
	app.writeJSON(w, http.StatusAccepted, title)
}
