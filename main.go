package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.POST("/deneme", Deneme)
	r.GET("/yazilar/:slug", Yazilar)
	http.ListenAndServe(":8080", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func Yazilar(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	data := params.ByName("slug")
	view.Execute(w, data)
}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("fileUpload")
	openFile, _ := os.OpenFile(header.Filename, os.O_CREATE, 066)
	io.Copy(openFile, file)
}
