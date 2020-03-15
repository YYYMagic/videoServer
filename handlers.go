package main

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

type HomePage struct {
	Name string
}

type LivePage struct {
	Name string
	Url string
	Type string
}

type VideoPage struct {
	Name string
	Url string
}


func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p := &HomePage{Name: "yyy"}
	t, err := template.ParseFiles("./urls/home.html")
	if err != nil {
		log.Printf("Parsing template home.html error: %s", err)
		return
	}
	t.Execute(w, p)
	return

}

func liveHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lid := params.ByName("lid")
	p := &LivePage{Name: "yyy"}
	t, err := template.ParseFiles("./urls/" + lid + ".html")
	if err != nil {
		log.Printf("Parsing template home.html error: %s", err)
		return
	}
	t.Execute(writer, p)
	return
}

func videoHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	vid := params.ByName("vid")
	p := &VideoPage{Name: videoMap[vid].Name, Url: videoMap[vid].Url}
	t, err := template.ParseFiles("./urls/videos.html")
	if err != nil {
		log.Printf("Parsing template home.html error: %s", err)
		return
	}
	t.Execute(writer, p)
	return
}


