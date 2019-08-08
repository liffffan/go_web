package main

import (
	"html/template"
	"net/http"
)

func main() {

	// NewServeMux 创建多路复用器，通过一些代码将接收到的请求重定向到处理器

	mux := http.NewServeMux()

	// 指定一个静态文件服务器
	files := http.FileServer(http.Dir("/public"))

	// StripPrefix 去除请求URL中的指定前缀，并把处理器传递给多路复用器
	mux.Handle("/static", http.StripPrefix("/static/", files))

	// 为了将发送至根URL的请求重定向到处理器，使用了 HandleFunc 函数
	// 这个函数接收一个URL和一个处理器的名称作为参数
	// 将针对给定URL的请求转发至指定的处理器进行处理，因此请求会被重定向到 index 函数
	// 所有处理器都会接收一个 ResponseWriter 和一个指向 Request 结构的指针作为参数
	// 并且所有请求的参数都可以通过 Request 结构体获取到，所以程序不需要显示给处理器传入任何请求参数
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// 创建处理器函数
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	if threads, err := data.Threads(); err == nil {
		templates.ExecuteTemplate(w, "layout", threads)

	}
}
