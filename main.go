package main

import(
		// "fmt"
		"net/http"
		"log"
		"project/handler"
		)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/",handler.HandlerIndex)
	mux.HandleFunc("/create-task",handler.CreatePage)
	mux.HandleFunc("/add",handler.InsertTask)
	mux.HandleFunc("/edit-task",handler.EditPage)
	mux.HandleFunc("/edit",handler.EditTask)
	mux.HandleFunc("/mark-done",handler.MarkDone)


	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	log.Println("Starting Port : 8000")

	err := http.ListenAndServe(":8000",mux)
	log.Fatal(err)
}

