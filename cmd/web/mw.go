package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// 1. 设置 header

		// 可跨域的源设置
		w.Header().Set("Access-Control-Allow-Origin", "https://*, http://*")
		// 设置支持的 HTTP 方法 
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS, UPDATE")
		// 设置请求可以携带的 headers 
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Token, session, Content-Length, X-CSRF-Token")
		// 获取非标准的 headers
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		// 是否发送 cookies
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// options 有效期
		w.Header().Set("Access-Control-Max-Age", "true")
	
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent) 
			return
		}

		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}


// 访问日志  
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		wc := &ResponseWithRecorder{
			ResponseWriter: w,
		}

		startTime := time.Now()
	
		// 除首页以外，移除所有请求路径后面的 "/"           
		if r.URL.Path != "/" {                                   
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")     
		}                                                        
                                                                 
		next.ServeHTTP(wc, r)                                    

		// 打日志
		log.Printf("| %3d | %13v | %15s | %s | %s |",
			wc.statusCode,          // 状态码
			time.Since(startTime),  // 耗费时间
			r.RemoteAddr,           // 请求 IP 地址
			r.Method,               // 请求方法
			r.RequestURI,           // 请求路由
		)	
	})
}



// 拷贝一份状态码
type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}
