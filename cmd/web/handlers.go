package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


type Credentials struct {
	Name       string `json:"name"        validate:"required,min=3,max=20"`
	Password   string `json:"password"    validate:"required,min=3,max=20,containsany=!@#$%*"`
	Email      string `json:"email"       validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,min=6,max=6"`
	Sex        int    `json:"sex"        validate:"omitempty"`
}

// 注册
func (app *App) register(w http.ResponseWriter, r *http.Request) {
	var req Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		return
	}

	arguments := ValidateHttpData(req)
	if len(arguments) > 0 {
		fmt.Println(arguments)
		return
	}
}

// 登录
func (app *App) login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	
	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	tokenString := base64.StdEncoding.EncodeToString([]byte(creds.Name))
	expiresAt := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiresAt,	
	})	
}


// 刷新
func (app *App) refresh(w http.ResponseWriter, r *http.Request) {

}


// 注销
func (app *App) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),  // 二选一，即可
		MaxAge:  -1,         
	})
}



func (app *App) fetchVideos(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token") 
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
	}
		
	// 参数路由，e.g. `/videos?size=10&&num=1`
	_ = r.URL.Query().Get("size")  // 每页包含的数据条目
	_ = r.URL.Query().Get("num")   // 当前页码

	// 其它操作
	_ = c

}


func (app *App) GetBook(w http.ResponseWriter, r *http.Request) {

	// 动态路由，e.g. `/books/the_golden_lotus`
	title := mux.Vars(r)["title"]
	
	app.writeJSON(w, http.StatusAccepted, title)
}


// 表单 Http.PostBody
// name := r.FormValue("username")
// pwd  := r.FormValue("password")
