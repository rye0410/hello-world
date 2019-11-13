package service

import (
	"net/http"

	"github.com/render"
)

func InfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "table", struct {
			ID       string
			Password string
		}{ID: req.Form["id"][0], Password: req.Form["password"][0]})
	}
}

//JsonHandler .
func JsonHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Title string `json:"title"`
		}{Title: "Hello Go !"})
	}
}

//UnknownHandler .
func UnknownHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "5xx Not found")
	}
}
